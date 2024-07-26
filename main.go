package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"poll-lite/models"
	"poll-lite/services/dataservice"
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
	writeResponseQuestionaryAsJson(&w, r, id)
}

func PollHandlerPost(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	var body models.AnswerBody
	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		fmt.Fprintf(w, "You've requested POST, but got an error back: %s at %s\n", r.URL.Path, time.Now())
	} else {
		dataservice.SaveAnswer(id, &body)
	}

	writeResponseAsJson(&w, id)
}

func ResultsHandlerGet(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	writeResponseResultAsJson(&w, r, id)
}

func writeResponseAsJson(w *http.ResponseWriter, val any) error {
	return json.NewEncoder(*w).Encode(val)
}

func writeNotFoundError(w *http.ResponseWriter, r *http.Request) {
	http.NotFound(*w, r)
}

func writeResponseResultAsJson(w *http.ResponseWriter, r *http.Request, id string) {
	// TODO: don't ignore errors, return proper response
	questionary, err := dataservice.GetQuestionary(id)

	if err != nil {
		writeNotFoundError(w, r)
		return
	}

	answers, err := prepareAnswersViewModel(id, questionary.IsAnonymous)

	if err != nil {
		writeNotFoundError(w, r)
		return
	}

	result := models.ResultViewModel{
		Questionary: questionary,
		Votes:       answers,
	}

	writeResponseAsJson(w, result)
}

func prepareAnswersViewModel(questionaryId string, isAnonymous bool) ([]models.AnswerBody, error) {
	answers, err := dataservice.GetAnswers(questionaryId)

	if err != nil {
		return []models.AnswerBody{}, err
	}

	if isAnonymous {
		for i := 0; i < len(answers); i++ {
			answer := &answers[i]
			answer.RespondentName = "N/A"
		}
	}
	return answers, nil
}

func writeResponseQuestionaryAsJson(w *http.ResponseWriter, r *http.Request, id string) {
	result, err := dataservice.GetQuestionary(id)

	if err != nil {
		writeNotFoundError(w, r)
		return
	}

	writeResponseAsJson(w, result)

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
