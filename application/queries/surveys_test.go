package queries_test

import (
	"context"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/golang/mock/gomock"
	"github.com/jbactad/loop/application/ports/mock"
	"github.com/jbactad/loop/application/queries"
	"github.com/jbactad/loop/domain"
	"github.com/stretchr/testify/assert"
)

func TestHandle_GetSurveys(t *testing.T) {
	testCases := []struct {
		name      string
		query     queries.GetSurveysQuery
		wantError assert.ErrorAssertionFunc
		setup     func(sp *mock.MockSurveyProvider)
		want      queries.GetSurveysQueryResponse
	}{
		{
			name: "with valid query, it should return surveys without error",
			query: queries.GetSurveysQuery{
				Limit: 10,
				Page:  0,
			},
			setup: func(sp *mock.MockSurveyProvider) {
				sp.EXPECT().GetSurveys(gomock.AssignableToTypeOf(context.Background()), 10, 0).Return(func() []*domain.Survey {
					surveys, err := SurveysTestData()
					if err != nil {
						t.Error(err)
					}

					return surveys
				}(), nil)
			},
			wantError: assert.NoError,
			want: queries.GetSurveysQueryResponse{
				Surveys: func() []*domain.Survey {
					surveys, err := SurveysTestData()
					if err != nil {
						t.Error(err)
					}

					return surveys
				}(),
			},
		},
		{
			name: "with invalid query, it should return error",
			query: queries.GetSurveysQuery{
				Limit: -1,
				Page:  -1,
			},
			setup: func(sp *mock.MockSurveyProvider) {
				sp.EXPECT().GetSurveys(gomock.Any(), gomock.Any(), gomock.Any()).Times(0)
			},
			wantError: assert.Error,
		},
		{
			name: "with error from provider, it should return error",
			query: queries.GetSurveysQuery{
				Limit: 10,
				Page:  0,
			},
			setup: func(sp *mock.MockSurveyProvider) {
				sp.EXPECT().GetSurveys(gomock.AssignableToTypeOf(context.Background()), 10, 0).
					Return(nil, assert.AnError)
			},
			wantError: assert.Error,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			surveyProvider := mock.NewMockSurveyProvider(ctrl)

			if tC.setup != nil {
				tC.setup(surveyProvider)
			}

			handler := queries.NewGetSurveysQueryHandler(surveyProvider)

			got, err := handler.Handle(context.Background(), tC.query)
			if tC.wantError(t, err) == true {
				return
			}

			assert.Equal(t, tC.want, got)
		})
	}
}

var surveysTestData []*domain.Survey

func SurveysTestData() (surveys []*domain.Survey, err error) {
	if surveysTestData != nil {
		return surveysTestData, nil
	}

	err = faker.FakeData(&surveys)
	if err != nil {
		return nil, err
	}
	surveysTestData = surveys

	return surveys, nil
}
