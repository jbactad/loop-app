// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package models

import (
	"fmt"
	"io"
	"strconv"
)

type NewSurvey struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Question    string `json:"question"`
}

type NewSurveyResponse struct {
	SurveyID string `json:"surveyId"`
	Answers  string `json:"answers"`
	Rating   int    `json:"rating"`
}

type Survey struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Question    string `json:"question"`
}

type SurveyResponse struct {
	ID     string  `json:"id"`
	Survey *Survey `json:"survey"`
	Answer string  `json:"answer"`
	Rating int     `json:"rating"`
}

type Role string

const (
	RoleManager  Role = "MANAGER"
	RoleEmployee Role = "EMPLOYEE"
)

var AllRole = []Role{
	RoleManager,
	RoleEmployee,
}

func (e Role) IsValid() bool {
	switch e {
	case RoleManager, RoleEmployee:
		return true
	}
	return false
}

func (e Role) String() string {
	return string(e)
}

func (e *Role) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Role(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Role", str)
	}
	return nil
}

func (e Role) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
