package ports

import (
	"context"
	"time"

	"github.com/jbactad/loop/domain"
)

//go:generate mockery --name=SurveyProvider --output=./mocks --outpkg=mocks --with-expecter
type SurveyProvider interface {
	GetSurveys(ctx context.Context, limit int, offset int) ([]*domain.Survey, error)
}

//go:generate mockery --name=SurveyCreator --output=./mocks --outpkg=mocks --with-expecter
type SurveyCreator interface {
	CreateSurvey(ctx context.Context, survey *domain.Survey) error
}

//go:generate mockery --name=UUIDGenerator --output=./mocks --outpkg=mocks --with-expecter
type UUIDGenerator interface {
	Generate() string
}

//go:generate mockery --name=TimeProvider --output=./mocks --outpkg=mocks --with-expecter
type TimeProvider interface {
	Now() time.Time
}
