package alexa

import (
	"encoding/json"
)

type CardType string

const (
	SimpleCardType      CardType = `Simple`
	StandardCardType    CardType = `Standard`
	LinkAccountCardType CardType = `LinkAccount`
)

var CardTypeDoesNotExist error

// OutputSpeech is an interface used to return the text to be spoken in the Response OutputSpeech and Reprompt fields.
//
// Types which implement this interface include PlainSpeech, a plain string containing the speech to render to the user,
// and SSMLSpeech, a string containing text marked up with SSML to render to the user.
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

type Card struct {
	Type          CardType
	Title         string
	Text          string
	LargeImageURL string
	SmallImageURL string
}

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

type Response struct {
	Version           string                 `json:"version"`
	SessionAttributes map[string]interface{} `json:"sessionAttributes"`
	Response          *ResponseData          `json:"response"`
}

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

type ResponseData struct {
	OutputSpeech     OutputSpeech `json:"outputSpeech,omitempty"`
	Card             *Card        `json:"card,omintempty"`
	Reprompt         OutputSpeech `json:"reprompt"`
	ShouldEndSession bool         `json:"shouldEndSession"`
}

type Directive struct{}
