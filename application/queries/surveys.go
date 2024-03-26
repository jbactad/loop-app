package queries

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

type ErrInvalidQuery struct {
	error
}

func (qs *Queries) GetSurveys(ctx context.Context, request GetSurveysQuery) (GetSurveysQueryResponse, error) {
	if (request.Limit < 0) || (request.Page < 0) {
		return GetSurveysQueryResponse{}, ErrInvalidQuery{}
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
		return GetSurveyByIdQueryResponse{}, ErrInvalidQuery{errors.New("id is required")}
	}

	survey, err := qs.repo.GetSurvey(ctx, request.Id)
	if err != nil {
		return GetSurveyByIdQueryResponse{}, err
	}

	return GetSurveyByIdQueryResponse{
		Survey: survey,
	}, nil
}
