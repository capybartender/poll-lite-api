package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	dataProvider "poll-lite/dataprovider"
	"poll-lite/models"
	"time"

	"github.com/rs/cors"
)

func CreateHandlerPost(w http.ResponseWriter, r *http.Request) {
	var body models.QuestionaryBody
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		fmt.Fprintf(w, "You've requested POST, but got an error back: %s at %s\n", r.URL.Path, time.Now())
	} else {
		newQuestionaryId := dataProvider.CreateQuestionary(body)
		json.NewEncoder(w).Encode(newQuestionaryId)
	}
}

func PollHandlerGet(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	writeResponseQuestionaryAsJson(&w, id)
}

func PollHandlerPost(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	var body models.Answer
	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		fmt.Fprintf(w, "You've requested POST, but got an error back: %s at %s\n", r.URL.Path, time.Now())
	} else {
		dataProvider.SaveAnswer(id, body)
	}

	writeResponseResultAsJson(&w, id)
}

func ResultsHandlerGet(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	writeResponseResultAsJson(&w, id)
}

func writeResponseResultAsJson(w *http.ResponseWriter, id string) {
	result := dataProvider.GetResult(id)
	json.NewEncoder(*w).Encode(result)
}

func writeResponseQuestionaryAsJson(w *http.ResponseWriter, id string) {
	result := dataProvider.GetQuestionary(id)
	json.NewEncoder(*w).Encode(result)
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
