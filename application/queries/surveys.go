package queries

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

var (
	ErrInvalidLimitOrOffset = errors.New("invalid limit or offset")
	ErrInvalidId            = errors.New("invalid id")
)

func (qs *Queries) GetSurveys(ctx context.Context, request GetSurveysQuery) (GetSurveysQueryResponse, error) {
	if (request.Limit < 0) || (request.Page < 0) {
		return GetSurveysQueryResponse{}, ErrInvalidLimitOrOffset
	}

	surveys, err := qs.repo.GetSurveys(ctx, request.Limit, request.Page)
	if err != nil {
		return GetSurveysQueryResponse{}, err
	}

	return GetSurveysQueryResponse{
		Surveys: surveys,
	}, nil
}

func (qs *Queries) GetSurveyByID(ctx context.Context, request GetSurveyByIdQuery) (GetSurveyByIdQueryResponse, error) {
	if uuid.Validate(request.Id) != nil {
		return GetSurveyByIdQueryResponse{}, ErrInvalidId
	}

	survey, err := qs.repo.GetSurvey(ctx, request.Id)
	if err != nil {
		return GetSurveyByIdQueryResponse{}, err
	}

	return GetSurveyByIdQueryResponse{
		Survey: survey,
	}, nil
}
