package dataprovider

import (
	"poll-lite/idgenerator"
	"poll-lite/models"
)

func GetResult(id string) models.ResultViewModel {
	result := models.ResultViewModel{
		Questionary: models.QuestionaryBody{
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
		},
		Votes: []models.AnswerBody{
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
		},
	}

	return result
}

func GetQuestionary(id string) models.QuestionaryBody {
	result := models.QuestionaryBody{
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

	return result
}

func SaveAnswer(id string, answer models.AnswerBody) {
	// TODO: save it
}

func CreateQuestionary(body models.QuestionaryBody) string {
	// TODO: save it
	return idgenerator.GenerateUniqueString()
}
