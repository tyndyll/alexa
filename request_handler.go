package alexa

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func NewAlexaRequestHandler(alexaRequestHandler func(*Request) (*Response, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		data, err := ioutil.ReadAll(req.Body)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		alexaRequest := &Request{}

		err = json.Unmarshal(data, alexaRequest)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		alexaResponse, err := alexaRequestHandler(alexaRequest)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		w.Header().Set(`Content-Type`, `application/json;charset=UTF-8`)
		data, err = json.Marshal(alexaResponse)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		w.Write(data)
	}
}