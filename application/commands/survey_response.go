package commands

import (
	"context"

	"github.com/jbactad/loop/domain"
)

type CreateSurveyResponseCommand struct {
	SurveyID string
	Answer   string
	Rating   int
}

func (cs *Commands) CreateSurveyResponse(ctx context.Context, cmd CreateSurveyResponseCommand) (*domain.SurveyResponse, error) {
	s, _ := cs.surveyCreatorProvider.GetSurvey(ctx, cmd.SurveyID)
	i := cs.uuidGenerator.Generate()
	n := cs.timeProvider.Now()
	sr := domain.NewSurveyResponse(i, s, cmd.Answer, cmd.Rating, n, n)

	_ = cs.surveyResponseCreatorProvider.CreateSurveyResponse(ctx, sr)

	return sr, nil
}
