package alexa

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http/httptest"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

const (
	demo_uid         = "urn:uuid:1335c695-cfb8-4ebb-abbd-80da344efa6b"
	demo_date        = "2016-05-23T22:34:51Z"
	demo_title       = "Amazon Developer Blog, week in review May 23rd"
	demo_text        = "Main information for the item"
	demo_stream      = "https://developer.amazon.com/public/community/blog/myaudiofile.mp3"
	demo_redirection = "https://developer.amazon.com/public/community/blog"

	demo_json = `{"uid":"%s","updateDate":"%s","titleText":"%s","mainText":"%s","streamUrl":"%s","redirectionUrl":"%s"}`
)

var expected_json = fmt.Sprintf(demo_json, demo_uid, demo_date, demo_title, demo_text,
	demo_stream, demo_redirection)

func TestFlashBriefingItem_MarshalJSON(t *testing.T) {
	Convey(`Given I have a FlashBriefingItem`, t, func() {
		date, err := time.Parse(time.RFC3339, demo_date)
		if err != nil {
			panic(err)
		}

		item := &FlashBriefingItem{
			ID:         demo_uid,
			Date:       date,
			Title:      demo_title,
			Text:       demo_text,
			AudioURL:   demo_stream,
			DisplayURL: demo_redirection,
		}

		Convey(`When I marshal the item to JSON`, func() {
			data, err := json.Marshal(item)

			Convey(`Then the error will be nil`, func() {
				So(err, ShouldBeNil)
			})

			Convey(`Then the correct json will be produced`, func() {
				So(string(data), ShouldEqual, expected_json)
			})
		})
	})

}

func TestFlashBriefing_MarshalJSON(t *testing.T) {
	Convey(`Given I have a FlashBriefingItem`, t, func() {
		date, err := time.Parse(time.RFC3339, demo_date)
		if err != nil {
			panic(err)
		}

		item := &FlashBriefingItem{
			ID:         demo_uid,
			Date:       date,
			Title:      demo_title,
			Text:       demo_text,
			AudioURL:   demo_stream,
			DisplayURL: demo_redirection,
		}

		Convey(`And I populate a FlashBriefing with that item`, func() {
			briefing := &FlashBriefing{
				Items: []*FlashBriefingItem{
					item,
					item,
				},
			}

			Convey(`When I marshal the briefing to JSON`, func() {
				data, err := json.Marshal(briefing)

				Convey(`Then the error will be nil`, func() {
					So(err, ShouldBeNil)
				})

				Convey(`Then the correct json will be produced`, func() {
					expected_item := fmt.Sprintf(demo_json, demo_uid, demo_date, demo_title, demo_text,
						demo_stream, demo_redirection)
					expected_json := fmt.Sprintf(`[%s,%s]`, expected_item, expected_item)
					So(string(data), ShouldEqual, expected_json)
				})
			})
		})
	})
}

func TestServeFlashBriefing(t *testing.T) {
	Convey(`Given I have a FlashBriefingItem`, t, func() {
		date, err := time.Parse(time.RFC3339, demo_date)
		if err != nil {
			panic(err)
		}

		item := &FlashBriefingItem{
			ID:         demo_uid,
			Date:       date,
			Title:      demo_title,
			Text:       demo_text,
			AudioURL:   demo_stream,
			DisplayURL: demo_redirection,
		}

		Convey(`And I populate a FlashBriefing with that item`, func() {
			briefing := &FlashBriefing{
				Items: []*FlashBriefingItem{
					item,
					item,
				},
			}

			Convey(`And I have a http handler`, func() {
				handler := FlashBriefingHandler(briefing)

				Convey(`And I have a HTTP request`, func() {
					req := httptest.NewRequest("GET", `/`, nil)

					Convey(`When I make the request`, func() {
						w := httptest.NewRecorder()

						handler.ServeHTTP(w, req)

						Convey(`Then the response will have the appropriate Content-Type`, func() {
							So(w.HeaderMap.Get(contentHeader), ShouldEqual, jsonContentType)
						})

						Convey(`Then the body will be the FlashBriefing as JSON`, func() {
							body, err := ioutil.ReadAll(w.Body)
							if err != nil {
								panic(err)
							}
							expectedJson, err := json.Marshal(briefing)
							if err != nil {
								panic(err)
							}
							So(body, ShouldResemble, expectedJson)
						})
					})

				})

			})
		})
	})
}
