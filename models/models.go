package models

import (
	"time"
)

type QuestionaryId string
type CreatorId string
type OptionId string

type CreatorName string
type CreatorSecret string
type CreatorSalt string

type Creator struct {
	Id   CreatorId
	Name string
}

type Option struct {
	Id   OptionId
	Text string
}

type Questionary struct {
	Id          QuestionaryId
	Title       string
	Description string

	CreatedBy        Creator
	CreatedAt        time.Time
	Options          []Option
	IsMultipleChoice bool
	IsAnonymous      bool
}

type Answer struct {
	QuestionaryId    QuestionaryId `json:"questionaryId"`
	RespondentName   string        `json:"respondentName"`
	RespondentChoice []OptionId    `json:"respondentChoice"`
	//CreatedAt        time.Time     `json:"createdAt"`
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
}

type ResultBody struct {
	Questionary QuestionaryBody `json:"questionary"`
	Votes       []Answer        `json:"votes"`
}
