package models

type QuestionaryId string
type CreatorId string
type OptionId string

type CreatorName string
type CreatorSecret string
type CreatorSalt string

type AnswerBody struct {
	QuestionaryId    QuestionaryId `json:"questionaryId"`
	RespondentName   string        `json:"respondentName"`
	RespondentChoice []OptionId    `json:"respondentChoice"`
}

/////////////////////////////////

type OptionBody struct {
	Id   string `json:"id"`
	Text string `json:"text"`
}

type QuestionaryBody struct {
	Id               string       `json:"id"`
	Title            string       `json:"title"`
	Description      string       `json:"description"`
	CreatedBy        string       `json:"createdBy"`
	Options          []OptionBody `json:"options"`
	IsMultipleChoice bool         `json:"isMultipleChoice"`
	IsAnonymous      bool         `json:"isAnonymous"`
	//CreatedAt        time.Time     `json:"createdAt"`
}

type ResultViewModel struct {
	Questionary QuestionaryBody `json:"questionary"`
	Votes       []AnswerBody    `json:"votes"`
}
