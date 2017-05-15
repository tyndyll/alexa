package alexa

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

const (
	demoUID         = "urn:uuid:1335c695-cfb8-4ebb-abbd-80da344efa6b"
	demoDate        = "2016-05-23T22:34:51Z"
	demoTitle       = "Amazon Developer Blog, week in review May 23rd"
	demoText        = "Main information for the item"
	demoStream      = "https://developer.amazon.com/public/community/blog/myaudiofile.mp3"
	demoRedirection = "https://developer.amazon.com/public/community/blog"
	demoJSON        = `{"uid":"%s","updateDate":"%s","titleText":"%s","mainText":"%s","streamUrl":"%s","redirectionUrl":"%s"}`
	max_items       = 5
)

var expectedJSON = fmt.Sprintf(demoJSON, demoUID, demoDate, demoTitle, demoText, demoStream, demoRedirection)

func TestFlashBriefingItem_MarshalJSON(t *testing.T) {
	Convey(`Given I have a FlashBriefingItem`, t, func() {
		date, err := time.Parse(time.RFC3339, demoDate)
		if err != nil {
			panic(err)
		}

		item := &FlashBriefingItem{
			ID:         demoUID,
			Date:       date,
			Title:      demoTitle,
			Text:       demoText,
			AudioURL:   demoStream,
			DisplayURL: demoRedirection,
		}

		Convey(`When I marshal the item to JSON`, func() {
			data, err := json.Marshal(item)

			Convey(`Then the error will be nil`, func() {
				So(err, ShouldBeNil)
			})

			Convey(`Then the correct json will be produced`, func() {
				So(string(data), ShouldEqual, expectedJSON)
			})
		})
	})
}

func TestFlashBriefing_MarshalJSON(t *testing.T) {
	Convey(`Given I have a FlashBriefingItem`, t, func() {
		date, err := time.Parse(time.RFC3339, demoDate)
		if err != nil {
			panic(err)
		}

		item := &FlashBriefingItem{
			ID:         demoUID,
			Date:       date,
			Title:      demoTitle,
			Text:       demoText,
			AudioURL:   demoStream,
			DisplayURL: demoRedirection,
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
					expectedBriefing := fmt.Sprintf(`[%s,%s]`, expectedJSON, expectedJSON)
					So(string(data), ShouldEqual, expectedBriefing)
				})
			})
		})
	})
}

func TestServeFlashBriefing(t *testing.T) {
	Convey(`Given I have a FlashBriefingItem`, t, func() {
		date, err := time.Parse(time.RFC3339, demoDate)
		if err != nil {
			panic(err)
		}

		item := &FlashBriefingItem{
			ID:         demoUID,
			Date:       date,
			Title:      demoTitle,
			Text:       demoText,
			AudioURL:   demoStream,
			DisplayURL: demoRedirection,
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
							marshaledJSON, err := json.Marshal(briefing)
							if err != nil {
								panic(err)
							}
							So(body, ShouldResemble, marshaledJSON)
						})
					})
				})
			})
		})
	})
}

func ExampleFlashBriefingHandler() {
	items := make([]*FlashBriefingItem, max_items)
	briefing := &FlashBriefing{
		Items: items,
	}

	for i := 0; i < max_items; i++ {
		briefing.Items[i] = &FlashBriefingItem{
			ID:    fmt.Sprintf(`id-%d`, i),
			Date:  time.Now(),
			Title: fmt.Sprintf(`Demo Entry %d`, i),
			Text:  fmt.Sprintf(`This is the entry for %d`, i),
		}
	}

	http.HandleFunc(`/`, FlashBriefingHandler(briefing))
	http.ListenAndServe(`:8080`, nil)
}
