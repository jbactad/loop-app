package commands

import (
	"context"

	"github.com/jbactad/loop/domain"
)

//go:generate mockery --name UseCases --output ./mocks --filename commands.go --with-expecter
type UseCases interface {
	CreateSurvey(ctx context.Context, cmd CreateSurveyCommand) (domain.Survey, error)
}
