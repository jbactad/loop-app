package queries_test

import (
	"context"
	"reflect"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/jbactad/loop/application/ports/mocks"
	"github.com/jbactad/loop/application/queries"
	"github.com/jbactad/loop/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestQueries_GetSurveys(t *testing.T) {
	var surveysTestData []*domain.Survey
	err := faker.FakeData(&surveysTestData)
	if err != nil {
		t.Error(err)
	}
	ctx := context.Background()
	tests := []struct {
		name    string
		query   queries.GetSurveysQuery
		want    queries.GetSurveysQueryResponse
		wantErr bool
		setup   func(sp *mocks.SurveyProvider)
	}{
		{
			name: "with valid query, it should return surveys without error",
			query: queries.GetSurveysQuery{
				Limit: 10,
				Page:  0,
			},
			setup: func(sp *mocks.SurveyProvider) {
				sp.EXPECT().
					GetSurveys(mock.IsType(context.Background()), 10, 0).
					Return(surveysTestData, nil)
			},
			want: queries.GetSurveysQueryResponse{
				Surveys: surveysTestData,
			},
			wantErr: false,
		},
		{
			name: "with invalid query, it should return error",
			query: queries.GetSurveysQuery{
				Limit: -1,
				Page:  -1,
			},
			wantErr: true,
		},
		{
			name: "with error from provider, it should return error",
			query: queries.GetSurveysQuery{
				Limit: 10,
				Page:  0,
			},
			setup: func(sp *mocks.SurveyProvider) {
				sp.EXPECT().
					GetSurveys(mock.IsType(context.Background()), 10, 0).
					Return(nil, assert.AnError)
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sp := mocks.NewSurveyProvider(t)
			if tt.setup != nil {
				tt.setup(sp)
			}

			qs := queries.New(sp)

			got, err := qs.GetSurveys(ctx, tt.query)

			if (err != nil) != tt.wantErr {
				t.Errorf("Queries.GetSurveys() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Queries.GetSurveys() = %v, want %v", got, tt.want)
			}
		})
	}
}
