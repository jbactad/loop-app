package ports

import (
	"context"

	"github.com/jbactad/loop/domain"
)

//go:generate mockery --name=SurveyProvider --output=./mocks --outpkg=mocks --with-expecter
type SurveyProvider interface {
	GetSurveys(ctx context.Context, limit int, offset int) ([]*domain.Survey, error)
}
