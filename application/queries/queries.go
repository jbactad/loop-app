package queries

import (
	"context"

	"github.com/jbactad/loop/application/ports"
)

//go:generate mockery --name UseCases --output ./mocks --filename queries.go --with-expecter
type UseCases interface {
	GetSurveys(ctx context.Context, request GetSurveysQuery) (GetSurveysQueryResponse, error)
}

type Queries struct {
	repo ports.SurveyProvider
}

func New(sp ports.SurveyProvider) *Queries {
	return &Queries{
		repo: sp,
	}
}
