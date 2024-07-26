package dataservice

import (
	"poll-lite/models"
	"poll-lite/services/idgenerator"
)

func GetQuestionary(id string) (models.QuestionaryBody, error) {

	questionary, err := getQuestionary(id)

	if err != nil {
		return models.QuestionaryBody{}, err
	}

	return questionary, nil
}

func GetAnswers(id string) ([]models.AnswerBody, error) {
	answers, err := getAnswers(id)

	if err != nil {
		return []models.AnswerBody{}, err
	}

	return answers, nil
}

func SaveAnswer(id string, answer *models.AnswerBody) error {
	return saveAnswer(id, answer)
}

func CreateQuestionary(body *models.QuestionaryBody) (string, error) {
	key, err := idgenerator.TakeNextUniqueKey()

	if err != nil {
		return "", err
	}

	err = saveQuestionary(key, body)

	if err != nil {
		return "", err
	}

	return key, nil
}
