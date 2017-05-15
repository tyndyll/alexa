package alexa

import (
	"encoding/json"
	"net/http"
	"time"
)

const (
	contentHeader   = `Content-Type`
	jsonContentType = `application/json`
)

// FlashBriefingItem is a representation of an item in an Alexa Flash Briefing.
type FlashBriefingItem struct {
	// Unique identifier for each feed item. UUID format preferred. This field is required for the feed
	ID string `json:"uid"`
	// Indicates freshness of feed item, and used to order items to be read or played from newest to oldest. Note
	// that older items may not be played or read. This field is required for the feed
	Date time.Time `json:"updateDate"`
	// The title of the feed item to display in the Alexa app. This field is required for the feed
	Title string `json:"titleText"`
	// For text feeds, the text that is read to the customer. This field is required for the feed, but may be an
	// empty string. Each feed item is currently limited to 4500 characters, and will be truncated if it exceeds
	// this length. The truncation will occur at the nearest full sentence below 4500 characters. The string should
	// be plain text and not contain special characters such as SSML, HTML or XML tags. The text should be properly
	// punctuated, short, and easily understood when read aloud. Commas (,) and semicolons (;) result in short
	// pauses. Periods (.), question marks (?) and exclamation points (!) result in longer pauses. Avoid using
	// non-standard punctuation as it could cause text to speech issues.
	Text string `json:"mainText,omitempty"`
	// HTTPS URL specifying the location of audio content for an audio feed. An audio item must contain a HTTPS URL
	// to audio content. The audio content should be 256kbps mono or stereo MP3. Content should not exceed 10
	// minutes in length. Program loudness should be -14 dB LUFS/LKFS. Alternatively, if not using LUFS or LKFS,
	// loudness should target a Total RMS value between -15 to -13 dB. The true-peak value should not exceed –2 dBFS
	AudioURL string `json:"streamUrl,omitempty"`
	// Provides the URL target for the Read More link in the Alexa app. This field is optional for the feed
	DisplayURL string `json:"redirectionUrl,omitempty"`
}

// A FlashBriefing provides audio or text content for a Flash Briefing skill. Alexa either plays or reads the feed
// contents to a customer.
type FlashBriefing struct {
	Items []*FlashBriefingItem
}

// MarshalJSON is a convenience function which manipulates the FlashBriefing into the JSON format expected by Alexa
func (briefing *FlashBriefing) MarshalJSON() ([]byte, error) {
	var interfaceSlice = make([]interface{}, len(briefing.Items))
	for i, item := range briefing.Items {
		interfaceSlice[i] = item
	}
	return json.Marshal(interfaceSlice)
}

// FlashBriefingHandler takes a pointer to a FlashBriefing and returns a HttpHandler which will respond to a HttpRequest
// with the FlashBriefing converted into JSON and appropriate response headers set
func FlashBriefingHandler(briefing *FlashBriefing) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Add(contentHeader, jsonContentType)
		if data, err := json.Marshal(briefing); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			w.Write(data)
		}
	})
}
