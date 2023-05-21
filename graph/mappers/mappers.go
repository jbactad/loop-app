package mappers

import (
	"github.com/jbactad/loop/domain"
	"github.com/jbactad/loop/graph/models"
)

func SurveysToSurveysResponse(surveys []*domain.Survey) []*models.Survey {
	response := make([]*models.Survey, len(surveys))
	for i, s := range surveys {
		response[i] = SurveyToResponse(s)
	}

	return response
}

func SurveyToResponse(survey *domain.Survey) *models.Survey {
	return &models.Survey{
		ID:          survey.ID(),
		Name:        survey.Name(),
		Question:    survey.Question(),
		Description: survey.Description(),
	}
}
