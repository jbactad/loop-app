package queries

import (
	"context"

	"github.com/jbactad/loop/application/ports"
	"github.com/jbactad/loop/domain"
)

type GetSurveysQuery struct {
	Limit int
	Page  int
}

type GetSurveysQueryResponse struct {
	Surveys []*domain.Survey
}
type GetSurveyByIdQuery struct {
	Id string
}

type GetSurveyByIdQueryResponse struct {
	Survey *domain.Survey
}

//go:generate mockery --name UseCases --output ./mocks --filename queries.go --with-expecter
type UseCases interface {
	GetSurveys(ctx context.Context, request GetSurveysQuery) (GetSurveysQueryResponse, error)
	GetSurveyByID(ctx context.Context, request GetSurveyByIdQuery) (GetSurveyByIdQueryResponse, error)
}

type Queries struct {
	repo ports.SurveyProvider
}

func New(sp ports.SurveyProvider) *Queries {
	return &Queries{
		repo: sp,
	}
}
