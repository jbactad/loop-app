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

func (cs *Commands) CreateSurvey(ctx context.Context, cmd CreateSurveyCommand) (domain.Survey, error) {
	now := cs.timeProvider.Now()
	id := cs.uuidGenerator.Generate()

	s := domain.NewSurvey(id, cmd.Name, cmd.Description, cmd.Question, now, now)

	err := cs.manager.CreateSurvey(ctx, s)
	if err != nil {
		return domain.Survey{}, err
	}

	return *s, nil
}
