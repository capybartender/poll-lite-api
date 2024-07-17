package dataservice

import (
	"errors"
	"poll-lite/models"
	"strconv"
)

//var questionaries = []models.QuestionaryBody{}

var questionariesRepository = make(map[string]models.QuestionaryBody, 0)

func getQuestionary(id string) (models.QuestionaryBody, error) {
	questionary, ok := questionariesRepository[id]
	if ok {
		return questionary, nil
	}

	return models.QuestionaryBody{}, errors.New("no questionary found")
}

func saveQuestionary(id string, body *models.QuestionaryBody) error {
	questionary := *body
	questionary.Id = id

	for i := 0; i < len(questionary.Options); i++ {
		option := &questionary.Options[i]
		option.Id = id + "_opt" + strconv.Itoa(i)
	}

	questionariesRepository[id] = questionary
	//questionaries = append(questionaries, questionary)
	return nil
}
