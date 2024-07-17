package dataservice

import (
	"poll-lite/models"
	"strconv"
)

var questionaries = []models.QuestionaryBody{}

func saveQuestionary(id string, body *models.QuestionaryBody) error {
	questionary := *body
	questionary.Id = id

	for i := 0; i < len(questionary.Options); i++ {
		option := &questionary.Options[i]
		option.Id = id + "_opt" + strconv.Itoa(i)
	}

	questionaries = append(questionaries, questionary)
	return nil
}
