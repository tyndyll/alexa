package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"

	"github.com/tyndyll/alexa"
)

func RollDice(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Error reading HTTP body:", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	req := &alexa.Request{}

	err = json.Unmarshal(data, req)
	if err != nil {
		log.Println("Error unmarshalling request:", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	var response string

	switch req.Detail.GetType() {
	case alexa.LaunchRequestType:
		response = `What dice would you like to roll`
	case alexa.IntentRequestType:
		response = fmt.Sprintf("You rolled a %d", rand.Intn(99)+1)
	}

	alexaResponse := &alexa.Response{
		Response: &alexa.ResponseData{
			OutputSpeech: alexa.PlainSpeech(response),
			Card: &alexa.Card{
				Type: alexa.SimpleCardType,
				Text: response,
			},
			ShouldEndSession: true,
		},
	}

	w.Header().Set(`Content-Type`, `application/json;charset=UTF-8`)
	data, err = json.Marshal(alexaResponse)
	if err != nil {
		log.Println(`Problem Marshalling Response: `, err)
	}
	w.Write(data)
}

func main() {
	//endpoint := alexa.RequestVerificationMiddleware(http.HandlerFunc(RollDice))
	endpoint := http.HandlerFunc(RollDice)
	http.Handle(`/`, endpoint)
	http.ListenAndServe(":9000", nil)
}
