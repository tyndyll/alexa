package alexa_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/tyndyll/alexa"
	"github.com/satori/go.uuid"
	"encoding/json"
)

func TestAudioPLayerRequest_GetType(t *testing.T) {
	Convey(`Given I have a AudioPlayerRequest instance`, t, func() {
		req := &alexa.AudioPlayerRequest{}

		Convey(`When I call GetType()`, func() {
			reqType := req.GetType()

			Convey(`Then the response will be the correct Audio Player Request type string`, func() {
				So(reqType, ShouldEqual)
			})
		})
	})
}

func TestAudioStream_UnmarshalJSON(t *testing.T) {
	Convey(`Given I have an AudioStream instance`, t, func() {
		stream := &alexa.AudioStream{
			URL: `http://example.org`,
			Token: uuid.NewV4().String(),
			ExpectedPreviousToken: uuid.NewV4().String(),
			Offset: 123456,
		}

		Convey(`When I marshal the instance to JSON`, func() {
			data, err := json.Marshal(stream)

			Convey(`Then the error will be nil`, func() {
				So(err, ShouldBeNil)
			})

			Convey(`Then the data will be blank`, func() {
				So(string(data), ShouldBeBlank)
			})
		})
	})
}