package alexa

/*
import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// RequestVerificationMiddleware verifies that a request was sent by Alexa
//
// Requests sent to your web service are transmitted over the Internet. To protect your endpoint from potential
// attackers, your web service should verify that incoming requests were sent by Alexa. Any requests coming from other
// sources should be rejected.
//
// There are two parts to validating incoming requests:
//   * Check the request signature to verify the authenticity of the request. Alexa signs all HTTPS requests.
//   * Check the request timestamp to ensure that the request is not an old request being sent as part of a “replay” attack.
func RequestVerificationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		body, _ := ioutil.ReadAll(req.Body)
		// Reset the request body

		req.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		request := &Request{}
		if err := json.Unmarshal(body, request); err != nil {
			panic(err)
		}

		if !TimestampInTolerance(request.Request.Timestamp) {
			http.Error(w, "Timestamp Not Recent", http.StatusBadRequest)
			return
		}

		if !VerifySignatureCertificateURL(req.Header.Get(certificateChainURL)) {
			http.Error(w, "Invalid Certificate Chain", http.StatusBadRequest)
			return
		}

		if !ValidateCertificate(req.Header.Get(certificateChainURL)) {
			http.Error(w, "Invalid Certificate", http.StatusBadRequest)
			return
		}

		next.ServeHTTP(w, req)
	})
}
*/
