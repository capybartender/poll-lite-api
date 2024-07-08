package dataProvider

import (
	"poll-lite/idgenerator"
	"poll-lite/models"
)

func GetResult(id string) models.ResultBody {
	result := models.ResultBody{
		Questionary: models.QuestionaryBody{
			Id:               id,
			Title:            "Test",
			Description:      "Test Questionary",
			MediaUrl:         "https://www.google.com/",
			CreatedBy:        "me",
			IsMultipleChoice: true,
			IsAnonymous:      true,
			Options: []models.OptionBody{
				{
					Id:       "1",
					Text:     "one",
					MediaUrl: "https://one.test",
				},
				{
					Id:       "2",
					Text:     "two",
					MediaUrl: "https://two.test",
				},
				{
					Id:       "3",
					Text:     "three",
					MediaUrl: "https://three.test",
				},
			},
		},
		Votes: []models.Answer{
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
		MediaUrl:         "https://www.google.com/",
		CreatedBy:        "me",
		IsMultipleChoice: true,
		IsAnonymous:      true,
		Options: []models.OptionBody{
			{
				Id:       "1",
				Text:     "one",
				MediaUrl: "https://one.test",
			},
			{
				Id:       "2",
				Text:     "two",
				MediaUrl: "https://two.test",
			},
			{
				Id:       "3",
				Text:     "three",
				MediaUrl: "https://three.test",
			},
		},
	}

	return result
}

func SaveAnswer(id string, answer models.Answer) {
	// TODO: save it
}

func CreateQuestionary(body models.QuestionaryBody) string {
	// TODO: save it
	return idgenerator.GenerateUniqueString()
}
