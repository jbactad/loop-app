package commands

import (
	"context"

	"github.com/jbactad/loop/domain"
)

type CreateSurveyCommand struct {
	Name        string
	Description string
	Question    string
}

func CreateSurvey(ctx context.Context, cmd CreateSurveyCommand) (domain.Survey, error) {
	return domain.Survey{}, nil
}
