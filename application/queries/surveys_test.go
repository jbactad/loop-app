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

func TestQueries_GetSurveyByID(t *testing.T) {
	ctx := context.Background()
	validSurvey := &domain.Survey{}
	faker.FakeData(validSurvey)
	type args struct {
		request queries.GetSurveyByIdQuery
	}
	tests := []struct {
		name    string
		args    args
		setup   func(sp *mocks.SurveyProvider)
		want    queries.GetSurveyByIdQueryResponse
		wantErr bool
	}{
		{
			name: "given valid query, it should return survey without error",
			args: args{
				request: queries.GetSurveyByIdQuery{
					// random uuid
					Id: "f4b3b3b3-4b3b-4b3b-4b3b-4b3b3b3b3b3b",
				},
			},
			setup: func(sp *mocks.SurveyProvider) {
				sp.EXPECT().GetSurvey(mock.IsType(context.Background()), "f4b3b3b3-4b3b-4b3b-4b3b-4b3b3b3b3b3b").
					Return(validSurvey, nil)
			},
			want: queries.GetSurveyByIdQueryResponse{
				Survey: validSurvey,
			},
			wantErr: false,
		},
		{
			name: "given invalid query, it should return error",
			args: args{
				request: queries.GetSurveyByIdQuery{
					Id: "",
				},
			},
			want:    queries.GetSurveyByIdQueryResponse{},
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
			got, err := qs.GetSurveyByID(ctx, tt.args.request)

			if (err != nil) != tt.wantErr {
				t.Errorf("Queries.GetSurveyById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.EqualValuesf(t, got, tt.want, "Queries.GetSurveyById() = %v, want %v", got, tt.want)
		})
	}
}
