package dataservice

import (
	"errors"
	"poll-lite/models"
	"strconv"
)

var questionariesRepository = make(map[string]models.QuestionaryBody, 0)

var answersRepository = make(map[string][]models.AnswerBody, 0)

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

	return nil
}

func getAnswers(id string) ([]models.AnswerBody, error) {
	answers, ok := answersRepository[id]
	if ok {
		return answers, nil
	}

	return nil, errors.New("no answers found")
}

func saveAnswer(id string, answer *models.AnswerBody) error {
	_, ok := answersRepository[id]
	if !ok {
		answersRepository[id] = make([]models.AnswerBody, 0)
	}
	answersRepository[id] = append(answersRepository[id], *answer)

	return nil
}
