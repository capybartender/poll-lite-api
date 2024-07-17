package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"poll-lite/dataservice"
	"poll-lite/models"
	"time"

	"github.com/rs/cors"
)

func CreateHandlerPost(w http.ResponseWriter, r *http.Request) {
	var body models.QuestionaryBody
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		// TODO: return proper error
		fmt.Fprintf(w, "You've requested POST, but got an error back: %s at %s\n", r.URL.Path, time.Now())
	} else {
		newQuestionaryId, err := dataservice.CreateQuestionary(&body)

		if err != nil {
			// TODO: return proper error
			fmt.Fprintf(w, "You've requested POST, but got an error back: %s at %s\n", r.URL.Path, time.Now())
		} else {
			writeResponseAsJson(&w, newQuestionaryId)
		}
	}
}

func PollHandlerGet(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	writeResponseQuestionaryAsJson(&w, id)
}

func PollHandlerPost(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	var body models.AnswerBody
	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		fmt.Fprintf(w, "You've requested POST, but got an error back: %s at %s\n", r.URL.Path, time.Now())
	} else {
		dataservice.SaveAnswer(id, body)
	}

	writeResponseResultAsJson(&w, id)
}

func ResultsHandlerGet(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	writeResponseResultAsJson(&w, id)
}

func writeResponseAsJson(w *http.ResponseWriter, val any) error {
	return json.NewEncoder(*w).Encode(val)
}

func writeResponseResultAsJson(w *http.ResponseWriter, id string) {
	result := dataservice.GetResult(id)
	writeResponseAsJson(w, result)
}

func writeResponseQuestionaryAsJson(w *http.ResponseWriter, id string) {
	result, err := dataservice.GetQuestionary(id)
	if err != nil {
		// TODO: return proper error
		fmt.Fprintf(*w, "You've requested POST, but got an error back at %v\n", time.Now())
	} else {
		writeResponseAsJson(w, result)
	}
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /create", CreateHandlerPost)

	mux.HandleFunc("GET /poll/{id}", PollHandlerGet)
	mux.HandleFunc("POST /poll/{id}", PollHandlerPost)

	mux.HandleFunc("GET /result/{id}", ResultsHandlerGet)

	handler := cors.Default().Handler(mux)
	http.ListenAndServe(":3080", handler)
}
