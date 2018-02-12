// DiceRoller is a simple custom skill that takes the number of dice and side, then generates a random number
//
// Dice roller is the most basic example of a simple customer Alexa skill. This example does not take advantage of the
// NewAlexaRequestHandler function in order to demonstrate the overall code flow, but this code essentially duplicates
// its functionality
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"flag"

	"github.com/tyndyll/alexa"
)

var portFlag int

const (
	launchMessage = `What dice would you like to roll`
)

func init() {
	flag.IntVar(&portFlag, `port`, 8000, `port to listen on`)
	flag.Parse()
}

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

	switch req.Request.GetType() {
	case alexa.LaunchRequestType:
		response = launchMessage
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

	log.Panicln(`Listening on port`, portFlag)
	if err := http.ListenAndServe(":9000", nil); err != nil {
		log.Panicln(`HTTP server stopped. Reason:`, err.Error())
	}
}
