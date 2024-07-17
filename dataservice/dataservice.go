package dataservice

import (
	"poll-lite/models"
	"poll-lite/services/idgenerator"
)

// TODO: remove
func mockQuestionary(id string) models.QuestionaryBody {
	return models.QuestionaryBody{
		Id:               id,
		Title:            "Test",
		Description:      "Test Questionary",
		CreatedBy:        "me",
		IsMultipleChoice: true,
		IsAnonymous:      true,
		Options: []models.OptionBody{
			{
				Id:   "1",
				Text: "one",
			},
			{
				Id:   "2",
				Text: "two",
			},
			{
				Id:   "3",
				Text: "three",
			},
		},
	}
}

// TODO: remove
func mockAnswers(id string) []models.AnswerBody {
	return []models.AnswerBody{
		{
			QuestionaryId:    models.QuestionaryId(id),
			RespondentName:   "me1",
			RespondentChoice: []models.OptionId{"1"},
		},
		{
			QuestionaryId:    models.QuestionaryId(id),
			RespondentName:   "me2",
			RespondentChoice: []models.OptionId{"1", "2"},
		},
		{
			QuestionaryId:    models.QuestionaryId(id),
			RespondentName:   "me3",
			RespondentChoice: []models.OptionId{"2", "3"},
		},
	}
}

func GetQuestionary(id string) (models.QuestionaryBody, error) {

	questionary, err := getQuestionary(id)

	if err != nil {
		return mockQuestionary(id), err
	}

	return questionary, nil
}

func GetAnswers(id string) ([]models.AnswerBody, error) {
	answers, err := getAnswers(id)

	if err != nil {
		return mockAnswers(id), err
	}

	return answers, nil
}

func SaveAnswer(id string, answer *models.AnswerBody) error {
	return saveAnswer(id, answer)
}

func CreateQuestionary(body *models.QuestionaryBody) (string, error) {
	// TODO: save it
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
