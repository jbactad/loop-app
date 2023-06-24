package commands

import (
	"context"
	"errors"

	"github.com/jbactad/loop/domain"
)

type CreateSurveyResponseCommand struct {
	SurveyID string
	Answer   string
	Rating   int
}

var ErrInvalidSurveyID = errors.New("invalid survey id")

func (cs *Commands) CreateSurveyResponse(ctx context.Context, cmd CreateSurveyResponseCommand) (*domain.SurveyResponse, error) {
	s, err := cs.surveyCreatorProvider.GetSurvey(ctx, cmd.SurveyID)
	if err != nil {
		return nil, ErrInvalidSurveyID
	}

	i := cs.uuidGenerator.Generate()
	n := cs.timeProvider.Now()
	sr := domain.NewSurveyResponse(i, s, cmd.Answer, cmd.Rating, n, n)

	_ = cs.surveyResponseCreatorProvider.CreateSurveyResponse(ctx, sr)

	return sr, nil
}
