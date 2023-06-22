package commands

import (
	"context"

	"github.com/jbactad/loop/application/ports"
	"github.com/jbactad/loop/domain"
)

//go:generate mockery --name UseCases --output ./mocks --filename commands.go --with-expecter
type UseCases interface {
	CreateSurvey(ctx context.Context, cmd CreateSurveyCommand) (*domain.Survey, error)
}

type Commands struct {
	manager       ports.SurveyCreator
	uuidGenerator ports.UUIDGenerator
	timeProvider  ports.TimeProvider
}

func New(creator ports.SurveyCreator, uuidGenerator ports.UUIDGenerator, timeProvider ports.TimeProvider) *Commands {
	return &Commands{
		manager:       creator,
		uuidGenerator: uuidGenerator,
		timeProvider:  timeProvider,
	}
}
