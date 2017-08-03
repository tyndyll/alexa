package alexa

import (
	//	"bytes"
	//	"encoding/json"
	//	"net/http"
	//	"net/http/httptest"
	"testing"
)

/*
func TestRequestVerificationMiddleware(t *testing.T) {
	Convey(`Given I have an Alexa endpoint`, t, func() {
		endpointCalled := false
		alexaEndpoint := http.HandlerFunc(func(http.ResponseWriter, *http.Request) {
			endpointCalled = true
		})

		Convey(`And I have an Alexa Request`, func() {
			alexaReq := &alexa.Request{
				Request: &alexa.LaunchRequest{},
			}

			Convey(`And the timestamp is greater the current time but within the tolerance`, func() {
				alexaReq.Request.Timestamp = time.Now().Add(alexa.TimestampVerificationTolerance / -2)

				Convey(`And I have a request`, func() {
					data, err := json.Marshal(alexaReq)
					if err != nil {
						panic(err)
					}
					body := bytes.NewBuffer(data)

					req, err := http.NewRequest(http.MethodGet, `/`, body)
					if err != nil {
						panic(err)
					}

					Convey(`When I make a request to the wrapped Alexa endpoint`, func() {
						wrapped := alexa.RequestVerificationMiddleware(alexaEndpoint)

						response := httptest.NewRecorder()

						wrapped.ServeHTTP(response, req)

						Convey(`Then the status code will be StatusOK`, func() {
							So(response.Code, ShouldEqual, http.StatusOK)
						})

						Convey(`Then the endpoint will have been called`, func() {
							So(endpointCalled, ShouldBeTrue)
						})
					})
				})
			})
			Convey(`And the timestamp is less the current time but within the tolerance`, func() {
				alexaReq.Request.Timestamp = time.Now().Add(alexa.TimestampVerificationTolerance / 2)

				Convey(`And I have a request`, func() {
					data, err := json.Marshal(alexaReq)
					if err != nil {
						panic(err)
					}
					body := bytes.NewBuffer(data)

					req, err := http.NewRequest(http.MethodGet, `/`, body)
					if err != nil {
						panic(err)
					}

					Convey(`When I make a request to the wrapped Alexa endpoint`, func() {
						wrapped := alexa.RequestVerificationMiddleware(alexaEndpoint)

						response := httptest.NewRecorder()

						wrapped.ServeHTTP(response, req)

						Convey(`Then the status code will be StatusOK`, func() {
							So(response.Code, ShouldEqual, http.StatusOK)
						})

						Convey(`Then the endpoint will have been called`, func() {
							So(endpointCalled, ShouldBeTrue)
						})
					})

				})
			})
			Convey(`And the timestamp is greater than the timestamp verification threshold`, func() {
				alexaReq.Request.Timestamp = time.Now().Add(alexa.TimestampVerificationTolerance * 2)

				Convey(`And I have a request`, func() {
					data, err := json.Marshal(alexaReq)
					if err != nil {
						panic(err)
					}
					body := bytes.NewBuffer(data)

					req, err := http.NewRequest(http.MethodGet, `/`, body)
					if err != nil {
						panic(err)
					}

					Convey(`When I make a request to the wrapped Alexa endpoint`, func() {
						wrapped := alexa.RequestVerificationMiddleware(alexaEndpoint)

						response := httptest.NewRecorder()

						wrapped.ServeHTTP(response, req)

						Convey(`Then the status code will be StatusBadRequest`, func() {
							So(response.Code, ShouldEqual, http.StatusBadRequest)
						})

						Convey(`Then the endpoint will have been called`, func() {
							So(endpointCalled, ShouldBeFalse)
						})
					})

				})
			})
			Convey(`And the timestamp is less than the timestamp verification threshold`, func() {
				alexaReq.Request.Timestamp = time.Now().Add(alexa.TimestampVerificationTolerance * -2)

				Convey(`And I have a request`, func() {
					data, err := json.Marshal(alexaReq)
					if err != nil {
						panic(err)
					}
					body := bytes.NewBuffer(data)

					req, err := http.NewRequest(http.MethodGet, `/`, body)
					if err != nil {
						panic(err)
					}

					Convey(`When I make a request to the wrapped Alexa endpoint`, func() {
						wrapped := alexa.RequestVerificationMiddleware(alexaEndpoint)

						response := httptest.NewRecorder()

						wrapped.ServeHTTP(response, req)

						Convey(`Then the status code will be StatusBadRequest`, func() {
							So(response.Code, ShouldEqual, http.StatusBadRequest)
						})

						Convey(`Then the endpoint will have been called`, func() {
							So(endpointCalled, ShouldBeFalse)
						})
					})

				})
			})
		})
	})
}*/

func TestValidateCertificate(t *testing.T) {
	//if !alexa.ValidateCertificate(`https://s3.amazonaws.com/echo.api/echo-api-cert.pem`) {
	//	t.Errorf("Invalid Certificate")
	//}
}
