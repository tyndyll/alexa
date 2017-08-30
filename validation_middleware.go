package alexa

import (
	"bytes"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	// TimestampVerificationTolerance is the +/- tolerance in nanoseconds that an incoming Alexa request may have before
	// it will be rejected by the RequestVerificationMiddleware
	TimestampVerificationTolerance time.Duration = 150000000000

	// certificateChainURL is the name of the header containing the request certificate
	certificateChainURL = "SignatureCertChainUrl"

	// signatureHeader is the name of the header on the request containing the encrypted signatureHeader.
	signatureHeader = "Signature"

	validAlternativeName = "echo-api.amazon.com"
)

// Package alexa requires that request are verified that they are by Alexa. Requests sent to your web service are
// transmitted over the Internet. To protect your endpoint from potential attackers, your web service should verify that
// incoming requests were sent by Alexa. Any requests coming from other sources should be rejected.
//
// There are two parts to validating incoming requests:
//   * Check the request signature to verify the authenticity of the request. Alexa signs all HTTPS requests.
//
//   This is required for certifying your Alexa skill and making it available to Amazon users. Web services that accept
//   unsigned requests or fail to verify the request signature are rejected.
//
//   * Check the request timestamp to ensure that the request is not an old request being sent as part of a "replay" attack
//
//   This is required for certifying your Alexa skill and making it available to Amazon users.
//
// Further information can be found at
// https://developer.amazon.com/public/solutions/alexa/alexa-skills-kit/docs/developing-an-alexa-skill-as-a-web-service

// RequestVerificationMiddleware verifies that a request was sent by Alexa
//
// Requests sent to your web service are transmitted over the Internet. To protect your endpoint from potential
// attackers, your web service should verify that incoming requests were sent by Alexa. Any requests coming from other
// sources should be rejected.
//
// There are two parts to validating incoming requests:
//   * Check the request signatureHeader to verify the authenticity of the request. Alexa signs all HTTPS requests.
//   * Check the request timestamp to ensure that the request is not an old request being sent as part of a “replay” attack.
func RequestVerificationMiddleware(next http.Handler) http.Handler {
	// TODO: cache certificates?
	// TODO: take application ID

	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			http.Error(w, "Cannot read body", http.StatusBadRequest)
			return
		}
		// Reset the request body
		req.Body = ioutil.NopCloser(bytes.NewBuffer(body))

		fmt.Println(string(body))
		request := &Request{}
		if err := json.Unmarshal(body, request); err != nil {
			http.Error(w, "err", http.StatusBadRequest)
			return
		}

		if !TimestampInTolerance(request.Detail.GetTimestamp()) {
			http.Error(w, "Timestamp Not Recent", http.StatusBadRequest)
			return
		}

		if !VerifySignatureCertificateURL(req.Header.Get(certificateChainURL)) {
			http.Error(w, "Invalid Certificate Chain", http.StatusBadRequest)
			return
		}

		if valid, err := ValidateCertificate(req.Header.Get(certificateChainURL), req.Header.Get(signatureHeader)); err != nil {
			http.Error(w, fmt.Sprintf("Invalid Certificate: %s", err.Error()), http.StatusInternalServerError)
		} else {
			if !valid {
				http.Error(w, "Invalid Certificate", http.StatusBadRequest)
				return
			}
		}

		next.ServeHTTP(w, req)
	})
}

// TimestampInTolerance takes a timestamp and verifies that it is valid within the bounds of the
// TimestampVerificationTolerance and it is not an old request being sent as part of a “replay” attack.
//
// This is required for certifying your Alexa skill and making it available to Amazon users.
func TimestampInTolerance(timestamp time.Time) bool {
	now := time.Now()
	// true if timestamp less than (currentTime + Tolerance) and timestamp > (currentTime - Tolerance)
	return timestamp.Unix() < now.Add(TimestampVerificationTolerance).Unix() && timestamp.Unix() > now.Add(-1*TimestampVerificationTolerance).Unix()
}

// VerifySignatureCertificateURL verifies the URL to ensure that it matches the format used by Amazon. This value can be
// found specified by the SignatureCertChainUrl header value on the request
//
// This is required for certifying your Alexa skill and making it available to Amazon users.
func VerifySignatureCertificateURL(certificateURL string) bool {
	parsed, err := url.Parse(certificateURL)
	if err != nil {
		return false
	}

	// TODO: Collapse and normalise the URL path

	// The protocol is equal to https (case insensitive).
	// The hostname is equal to s3.amazonaws.com (case insensitive).
	// The path starts with /echo.api/ (case sensitive).
	// If a port is defined in the URL, the port is equal to 443.
	if parsed.Scheme != `https` || parsed.Hostname() != `s3.amazonaws.com` ||
		strings.Index(parsed.Path, `echo.api`) != 1 || parsed.Port() != "" && parsed.Port() != "443" {
		return false
	}
	return true
}

// ValidateCertificate verifies the signatureHeader in the HTTP headers
func ValidateCertificate(certificateURL, signature string) (bool, error) {
	// Download the PEM-encoded X.509 certificate chain that Alexa used to sign the message as specified by the
	// certificateURL header value on the request.
	certResponse, err := http.Get(certificateURL)
	if err != nil {
		return false, fmt.Errorf("Cannot download certificate")
	}

	certBody, err := ioutil.ReadAll(certResponse.Body)
	if err != nil {
		panic(err)
	}

	// TODO: Handle the unused variable
	block, _ := pem.Decode(certBody)
	if block == nil {
		return false, fmt.Errorf(`No PEM data found`)
	}
	certs, err := x509.ParseCertificates(block.Bytes)
	if err != nil {
		return false, err
	}

	// TODO: Remove
	// log.Println("TOTAL NUMBER OF CERTS: ", len(certs))
	for _, cert := range certs {
		// TODO:
		// The signing certificate has not expired (examine both the Not Before and Not After dates)
		// The domain echo-api.amazon.com is present in the Subject Alternative Names (SANs) section of the signing certificate
		// All certificates in the chain combine to create a chain of trust to a trusted root CA certificate
		if err := cert.VerifyHostname(validAlternativeName); err != nil {
			return false, fmt.Errorf("Could not validate name %s: %s", validAlternativeName, err)
		}

		// Base64-decode the Signature header value on the request to obtain the encrypted signatureHeader.
		//encryptedSig, err := base64.StdEncoding.DecodeString(signature)
		_, err := base64.StdEncoding.DecodeString(signature)
		if err != nil {
			// TODO: Remove
			// fg
			// log.Println("Decrypt : ", err)
			return false, err
		}

		// Use the public key extracted from the signing certificate to decrypt the encrypted signatureHeader to produce
		// the asserted hash value.
		// fmt.Println("Encryption key: ", cert.PublicKeyAlgorithm)
		//publicKey := cert.PublicKey

	}

	// Generate a SHA-1 hash value from the full HTTPS request body to produce the derived hash value

	// Compare the asserted hash value and derived hash values to ensure that they match.
	return true, nil
}
