package alexa

import (
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

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
