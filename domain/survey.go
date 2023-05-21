package domain

type Survey struct {
	id              string
	name            string
	description     string
	question        string
	surveyResponses []*SurveyResponse
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

func (s *Survey) SurveyResponses() []*SurveyResponse {
	return s.surveyResponses
}

func (s *Survey) SetName(name string) {
	s.name = name
}

func (s *Survey) SetDescription(description string) {
	s.description = description
}

func (s *Survey) SetQuestion(question string) {
	s.question = question
}

func (s *Survey) AddSurveyResponse(surveyResponse *SurveyResponse) {
	s.surveyResponses = append(s.surveyResponses, surveyResponse)
}

type SurveyResponse struct {
	id     string
	survey *Survey
	answer string
	rating int
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
