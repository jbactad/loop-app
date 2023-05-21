package queries

import (
	"context"

	"github.com/jbactad/loop/application/ports"
	"github.com/jbactad/loop/domain"
)

type ErrInvalidQuery struct {
	error
}

type GetSurveysQuery struct {
	Limit int
	Page  int
}

type GetSurveysQueryResponse struct {
	Surveys []*domain.Survey
}

type GetSurveysQueryHandler struct {
	repo ports.SurveyProvider
}

func NewGetSurveysQueryHandler(sp ports.SurveyProvider) *GetSurveysQueryHandler {
	return &GetSurveysQueryHandler{
		repo: sp,
	}
}

func (h *GetSurveysQueryHandler) Handle(ctx context.Context, request GetSurveysQuery) (GetSurveysQueryResponse, error) {
	if (request.Limit < 0) || (request.Page < 0) {
		return GetSurveysQueryResponse{}, ErrInvalidQuery{}
	}

	surveys, err := h.repo.GetSurveys(ctx, request.Limit, request.Page)
	if err != nil {
		return GetSurveysQueryResponse{}, err
	}

	return GetSurveysQueryResponse{
		Surveys: surveys,
	}, nil
}
