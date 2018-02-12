package alexa

import (
	"encoding/json"
)

type CardType string

const (
	// SimpleCardType is a card that contains a title and plain text content.
	SimpleCardType CardType = `Simple`
	// StandardCardType is a card that contains a title, text content, and an image to display.
	StandardCardType CardType = `Standard`
	// LinkAccountCardType is a card that displays a link to an authorization URL that the user can use to link their
	// Alexa account with a user in another system
	LinkAccountCardType CardType = `LinkAccount`
)

// CardTypeDoesNotExist is an error returned when the output card has been set to an unknown CardType
var CardTypeDoesNotExist error

// OutputSpeech is an interface used to return the text to be spoken in the Response OutputSpeech and Reprompt fields.
//
// Types which implement this interface include PlainSpeech, a plain string containing the speech to render to the user,
// and SSMLSpeech, a string containing text marked up with SSML to render to the user.
//
type OutputSpeech interface {
	MarshalJSON() ([]byte, error)
}

// PlainSpeech is string containing the text to be spoken, rendered as PlainText in the response.
type PlainSpeech string

// MarshalJSON converts the PlainSpeech string into a JSON object, setting the type field and naming the output field
// appropriately
func (t PlainSpeech) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type string `json:"type"`
		Text string `json:"text"`
	}{"PlainText", string(t)})
}

// SSMLSpeech is a string containing text marked up with SSML to be spoken. The type performs no validation on the
// passed SSML
type SSMLSpeech string

// MarshalJSON converts the SSMLSpeech string into a JSON object, setting the type field and naming the output field
// appropriately
func (t SSMLSpeech) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type string `json:"type"`
		SSML string `json:"ssml"`
	}{"SSML", string(t)})
}

// Card describes a visual interface that will be presented to the user.
//
// A card can only be included when sending a response to a LaunchRequest or IntentRequest. All of the text included in
// a card cannot exceed 8000 characters. This includes the title, content, text, and image URLs.
//
// An image URL (smallImageUrl or largeImageUrl) cannot exceed 2000 characters.
type Card struct {
	// Type describes the type of card to render
	Type CardType

	// Title is a string containing the title of the card.
	Title string

	// Text is a string containing the text content. Note that this field represents "content" in a SimpleCardType and
	// "text" in  a StandardCardType
	Text string

	// LargeImageURL is a string that specifies the URLs for a large image to display on a Standard card.
	LargeImageURL string

	// SmallImageURL is a string that specifies the URLs for a small image to display on a Standard card.
	SmallImageURL string
}

// MarshalJSON implements the json.Marshaler interface for the Card type. It selects the appropriate CardType and
// populates the appropriate field
func (card *Card) MarshalJSON() (data []byte, err error) {
	switch card.Type {
	case SimpleCardType:
		data, err = json.Marshal(&struct {
			Type    string `json:"type"`
			Title   string `json:"title"`
			Content string `json:"content,omitempty"`
		}{
			string(SimpleCardType),
			card.Title,
			card.Text,
		})
	case StandardCardType:
		type Images struct {
			Large string `json:"largeImageUrl,omitempty"`
			Small string `json:"smallImageUrl,omitempty"`
		}

		data, err = json.Marshal(&struct {
			Type  string  `json:"type"`
			Title string  `json:"title"`
			Text  string  `json:"text,omitempty"`
			Image *Images `json:"image,omitempty"`
		}{
			string(StandardCardType),
			card.Title,
			card.Text,
			&Images{card.LargeImageURL, card.SmallImageURL},
		})
	case LinkAccountCardType:
		data, err = json.Marshal(&struct {
			Type string `json:"type"`
		}{string(LinkAccountCardType)})
	default:
		err = CardTypeDoesNotExist
	}
	return data, err
}

// Response is the object that will be returned from a user request
type Response struct {
	// SessionAttributes is a map of key-value pairs to persist in the session
	// Session attributes are ignored by the Alexa service if they are included on a response to an AudioPlayer or
	// a PlaybackController request.
	SessionAttributes map[string]interface{} `json:"sessionAttributes"`

	// Response defines what to render to the user and whether to end the current session
	Response *ResponseData `json:"response"`
}

// MarshalJSON implements the json.Marshaler interface for the Response type. It selects the appropriate Response type
// and populates the version field
//
// The total size of your response cannot exceed 24 kilobytes
func (response *Response) MarshalJSON() (data []byte, err error) {
	type Alias Response
	return json.Marshal(&struct {
		Version string `json:"version"`
		*Alias
	}{
		Version: "1.0",
		Alias:   (*Alias)(response),
	})
}

// ResponseData contains the content that will be presented to a user
//
// If you include both standard properties and an AudioPlayer directive, Alexa processes the standard properties first.
// For example, if you provide OutputSpeech in the same response as an Play directive, Alexa speaks the provided text
// before beginning to stream the audio.
type ResponseData struct {
	// OutputSpeech contains the speech to render to the user. The OutputSpeech response cannot exceed 8000 characters.
	OutputSpeech OutputSpeech `json:"outputSpeech,omitempty"`

	// Card contains a card to render to the Amazon Alexa App. All of the text included in a card cannot exceed 8000
	// characters. This includes the title, content, text, and image URLs.
	Card *Card `json:"card,omintempty"`

	// Reprompt contains the output speech that will be provided if a re-prompt is necessary. This is used if the the
	// service keeps the session open after sending the response, but the user does not respond with anything that maps
	// to an intent defined in your voice interface while the audio stream is open.
	Reprompt OutputSpeech `json:"reprompt"`

	// ShouldEndSession is boolean value with true meaning that the session should end after Alexa speaks the response,
	// or false if the session should remain active.
	ShouldEndSession bool `json:"shouldEndSession"`

	// Directives is an array of directives specifying device-level actions to take using a particular interface
	Directives []Directive `json:"directives"`
}

// NewPlainSpeechResponse is a utility function that takes the Output Speech to be delivered in the response, populates
// it in a Response and then marshals that response into JSON
func NewPlainSpeechResponse(outputSpeech string) ([]byte, error) {
	response := &Response{
		Response: &ResponseData{
			OutputSpeech:     PlainSpeech(outputSpeech),
			ShouldEndSession: true,
		},
	}
	return json.Marshal(response)
}
