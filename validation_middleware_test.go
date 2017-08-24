package alexa

import (
	//	"bytes"
	//	"encoding/json"
	//	"net/http"
	//	"net/http/httptest"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
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

func TestTimestampInTolerance(t *testing.T) {
	Convey(`Given I have a timestamp that is greater than the current time but within the tolerance`, t, func() {
		timestamp := time.Now().Add(TimestampVerificationTolerance / 2)

		Convey(`When I call TimestampInTolerance`, func() {
			result := TimestampInTolerance(timestamp)

			Convey(`Then the result will be true`, func() {
				So(result, ShouldBeTrue)
			})
		})
	})

	Convey(`Given I have a timestamp that is less than the current time but within the tolerance`, t, func() {
		timestamp := time.Now().Add(TimestampVerificationTolerance / -2)

		Convey(`When I call TimestampInTolerance`, func() {
			result := TimestampInTolerance(timestamp)

			Convey(`Then the result will be true`, func() {
				So(result, ShouldBeTrue)
			})
		})
	})

	Convey(`Given I have a timestamp that is greater than the current time but not within the tolerance`, t, func() {
		timestamp := time.Now().Add(TimestampVerificationTolerance * 2)

		Convey(`When I call TimestampInTolerance`, func() {
			result := TimestampInTolerance(timestamp)

			Convey(`Then the result will be false`, func() {
				So(result, ShouldBeFalse)
			})
		})
	})

	Convey(`Given I have a timestamp that is less than the current time but not within the tolerance`, t, func() {
		timestamp := time.Now().Add(TimestampVerificationTolerance * -2)

		Convey(`When I call TimestampInTolerance`, func() {
			result := TimestampInTolerance(timestamp)

			Convey(`Then the result will be false`, func() {
				So(result, ShouldBeFalse)
			})
		})
	})
}

func TestVerifySignatureCertificateURL(t *testing.T) {
	Convey(`Given I have a valid certificate URL`, t, func() {
		path := `https://s3.amazonaws.com/echo.api/echo-api-cert.pem`

		Convey(`When I call VerifySignatureCertificateURL`, func() {
			result := VerifySignatureCertificateURL(path)

			Convey(`Then the result will be true`, func() {
				So(result, ShouldBeTrue)
			})
		})
	})

	Convey(`Given I have a valid certificate URL containing a port`, t, func() {
		path := `https://s3.amazonaws.com:443/echo.api/echo-api-cert.pem`

		Convey(`When I call VerifySignatureCertificateURL`, func() {
			result := VerifySignatureCertificateURL(path)

			Convey(`Then the result will be true`, func() {
				So(result, ShouldBeTrue)
			})
		})
	})

	Convey(`Given I have a valid certificate URL containing a valid un-normalized path`, t, func() {
		path := `https://s3.amazonaws.com/echo.api/../echo.api/echo-api-cert.pem`

		Convey(`When I call VerifySignatureCertificateURL`, func() {
			result := VerifySignatureCertificateURL(path)

			Convey(`Then the result will be true`, func() {
				So(result, ShouldBeTrue)
			})
		})
	})

	Convey(`Given I have a certificate URL containing an invalid scheme`, t, func() {
		path := `http://s3.amazonaws.com/echo.api/echo-api-cert.pem`

		Convey(`When I call VerifySignatureCertificateURL`, func() {
			result := VerifySignatureCertificateURL(path)

			Convey(`Then the result will be false`, func() {
				So(result, ShouldBeFalse)
			})
		})
	})

	Convey(`Given I have a certificate URL containing an invalid hostname`, t, func() {
		path := `https://notamazon.com/echo.api/echo-api-cert.pem`

		Convey(`When I call VerifySignatureCertificateURL`, func() {
			result := VerifySignatureCertificateURL(path)

			Convey(`Then the result will be false`, func() {
				So(result, ShouldBeFalse)
			})
		})
	})

	Convey(`Given I have a certificate URL containing an badly cased path`, t, func() {
		path := `https://s3.amazonaws.com/EcHo.aPi/echo-api-cert.pem`

		Convey(`When I call VerifySignatureCertificateURL`, func() {
			result := VerifySignatureCertificateURL(path)

			Convey(`Then the result will be false`, func() {
				So(result, ShouldBeFalse)
			})
		})
	})

	Convey(`Given I have a certificate URL containing an invalid path`, t, func() {
		path := `https://s3.amazonaws.com/invalid.path/echo-api-cert.pem`

		Convey(`When I call VerifySignatureCertificateURL`, func() {
			result := VerifySignatureCertificateURL(path)

			Convey(`Then the result will be false`, func() {
				So(result, ShouldBeFalse)
			})
		})
	})

	Convey(`Given I have a certificate URL containing an invalid port`, t, func() {
		path := `https://s3.amazonaws.com:563/echo.api/echo-api-cert.pem`

		Convey(`When I call VerifySignatureCertificateURL`, func() {
			result := VerifySignatureCertificateURL(path)

			Convey(`Then the result will be false`, func() {
				So(result, ShouldBeFalse)
			})
		})
	})

	Convey(`Given I have an invalid URL`, t, func() {
		path := `definitely not a URL`

		Convey(`When I call VerifySignatureCertificateURL`, func() {
			result := VerifySignatureCertificateURL(path)

			Convey(`Then the result will be false`, func() {
				So(result, ShouldBeFalse)
			})
		})
	})
}
