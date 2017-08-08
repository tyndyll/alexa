package alexa

import (
	"crypto/x509"
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

	certificateChainURL  = "SignatureCertChainUrl"
	validAlternativeName = "echo-api.amazon.com"
)

// TimestampInTolerance takes a timestamp and verifies that it is valid within the bounds of the
// TimestampVerificationTolerance
func TimestampInTolerance(timestamp time.Time) bool {
	now := time.Now()
	// true if timestamp less than (currentTime + Tolerance) and timestamp > (currentTime - Tolerance)
	return timestamp.Unix() < now.Add(TimestampVerificationTolerance).Unix() && timestamp.Unix() > now.Add(-1*TimestampVerificationTolerance).Unix()
}

func VerifySignatureCertificateURL(certificateURL string) bool {
	parsed, err := url.Parse(certificateURL)
	if err != nil {
		return false
	}

	if parsed.Scheme != `https` || parsed.Hostname() != `s3.amazonaws.com` ||
		strings.Index(parsed.Path, `echo.api`) != 1 || parsed.Port() != "" && parsed.Port() != "443" {
		return false
	}
	return true
}

func ValidateCertificate(certificateURL string) bool {
	certResponse, err := http.Get(certificateURL)
	if err != nil {
		fmt.Println("Cannot download certificate")
		return false
	}

	certBody, err := ioutil.ReadAll(certResponse.Body)
	if err != nil {
		panic(err)
	}

	block, _ := pem.Decode(certBody)
	//	if block == nil || block.Type != "PUBLIC KEY" {
	//		log.Println("failed to decode PEM block containing public key")
	//		return false
	//	}

	certs, err := x509.ParseCertificates(block.Bytes)
	if err != nil {
		fmt.Printf(err.Error())
		return false
	}

	for _, cert := range certs {
		if err := cert.VerifyHostname(validAlternativeName); err != nil {
			fmt.Printf("Could not validate name %s: %s", validAlternativeName, err)
			return false
		}

		fmt.Printf("Cert valid not before %+v and not after %+v", cert.NotBefore, cert.NotAfter)
	}

	return true
}
