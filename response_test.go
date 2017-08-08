package alexa

import (
	"encoding/json"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPlainSpeech_MarshalJSON(t *testing.T) {
	Convey(`Given I have a PlainSpeech instance`, t, func() {
		var t PlainSpeech = `Test "Text"`

		Convey(`When I marshal it to JSON`, func() {
			output, err := json.Marshal(t)
			if err != nil {
				panic(err)
			}

			Convey(`Then the resulting JSON will be correct`, func() {
				So(string(output), ShouldEqual, `{"type":"PlainText","text":"Test \"Text\""}`)
			})
		})
	})
}

func TestSSMLSpeech_MarshalJSON(t *testing.T) {
	Convey(`Given I have a SSMLSpeech instance`, t, func() {
		var t SSMLSpeech = `Test "Text"`

		Convey(`When I marshal it to JSON`, func() {
			output, err := json.Marshal(t)
			if err != nil {
				panic(err)
			}

			Convey(`Then the resulting JSON will be correct`, func() {
				So(string(output), ShouldEqual, `{"type":"SSML","ssml":"Test \"Text\""}`)
			})
		})
	})
}
