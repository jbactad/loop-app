package domain

import "time"

type Survey struct {
	id          string
	name        string
	description string
	question    string
	createdAt   time.Time
	updatedAt   time.Time
}

func NewSurvey(
	id string,
	name string,
	description string,
	question string,
	createdAt time.Time,
	updatedAt time.Time,
) *Survey {
	return &Survey{
		id:          id,
		name:        name,
		description: description,
		question:    question,
		createdAt:   createdAt,
		updatedAt:   updatedAt,
	}
}

func (s *Survey) ID() string {
	return s.id
}

func (s *Survey) Name() string {
	return s.name
}

func (s *Survey) Description() string {
	return s.description
}

func (s *Survey) Question() string {
	return s.question
}

func (s *Survey) CreatedAt() time.Time {
	return s.createdAt
}

func (s *Survey) UpdatedAt() time.Time {
	return s.updatedAt
}

type SurveyResponse struct {
	id        string
	survey    *Survey
	answer    string
	rating    int
	createdAt time.Time
	updatedAt time.Time
}

func NewSurveyResponse(
	id string,
	survey *Survey,
	answer string,
	rating int,
	createdAt time.Time,
	updatedAt time.Time,
) *SurveyResponse {
	return &SurveyResponse{
		id:        id,
		survey:    survey,
		answer:    answer,
		rating:    rating,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}
}

func (s *SurveyResponse) ID() string {
	return s.id
}

func (s *SurveyResponse) Survey() *Survey {
	return s.survey
}

func (s *SurveyResponse) Answer() string {
	return s.answer
}

func (s *SurveyResponse) Rating() int {
	return s.rating
}

func (s *SurveyResponse) CreatedAt() time.Time {
	return s.createdAt
}

func (s *SurveyResponse) UpdatedAt() time.Time {
	return s.updatedAt
}
