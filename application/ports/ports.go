package ports

import (
	"context"

	"github.com/jbactad/loop/domain"
)

//go:generate mockgen -source=../ports/ports.go -destination=../ports/mock/mock.go -package=mock

type SurveyProvider interface {
	GetSurveys(ctx context.Context, limit int, offset int) ([]*domain.Survey, error)
}
