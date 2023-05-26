package resolvers_test

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/jbactad/loop/application/queries"
	"github.com/jbactad/loop/application/queries/mocks"
	"github.com/jbactad/loop/domain"
	"github.com/jbactad/loop/graph/models"
	"github.com/jbactad/loop/graph/resolvers"
	"github.com/stretchr/testify/mock"
)

func Test_queryResolver_Surveys(t *testing.T) {
	var surveysTestData []*domain.Survey

	err := faker.FakeData(&surveysTestData)
	if err != nil {
		t.Error(err)
	}
	ctx := context.Background()
	type args struct {
		limit *int
		page  *int
	}
	tests := []struct {
		name    string
		args    args
		want    []*models.Survey
		wantErr bool
		setup   func(md *mocks.UseCases)
	}{
		{
			name: "given meidatr returns a list of surveys, then return a list of surveys",
			setup: func(md *mocks.UseCases) {
				resp := queries.GetSurveysQueryResponse{
					Surveys: surveysTestData,
				}

				md.EXPECT().GetSurveys(ctx, queries.GetSurveysQuery{
					Limit: 10,
					Page:  0,
				}).
					Return(resp, nil).
					Once()
			},
			args: args{
				limit: func() *int {
					limit := 10
					return &limit
				}(),
				page: func() *int {
					page := 0
					return &page
				}(),
			},
			want: func() []*models.Survey {
				var modelsSurveys []*models.Survey
				for _, s := range surveysTestData {
					modelsSurveys = append(modelsSurveys, &models.Survey{
						ID:          s.ID(),
						Name:        s.Name(),
						Description: s.Description(),
						Question:    s.Question(),
						CreatedAt:   s.CreatedAt(),
						UpdatedAt:   s.UpdatedAt(),
					})
				}
				return modelsSurveys
			}(),
		},
		{
			name: "given no limit and page, then it should default to 10 and 0 respectively",
			setup: func(md *mocks.UseCases) {
				md.EXPECT().GetSurveys(ctx, queries.GetSurveysQuery{
					Limit: 10,
					Page:  0,
				}).
					Return(queries.GetSurveysQueryResponse{}, nil).
					Once()
			},
			args: args{
				limit: nil,
				page:  nil,
			},
			want: func() []*models.Survey {
				return []*models.Survey{}
			}(),
		},
		{
			name: "given mediatr returned error, then return error",
			setup: func(md *mocks.UseCases) {
				md.EXPECT().GetSurveys(ctx, mock.IsType(queries.GetSurveysQuery{})).
					Return(queries.GetSurveysQueryResponse{}, errors.New("error happened")).
					Once()
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			qu := mocks.NewUseCases(t)
			if tt.setup != nil {
				tt.setup(qu)
			}
			rs := &resolvers.Resolver{
				Queries: qu,
			}
			r := rs.Query()

			got, err := r.Surveys(ctx, tt.args.limit, tt.args.page)

			if (err != nil) != tt.wantErr {
				t.Errorf("queryResolver.Surveys() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("queryResolver.Surveys() = %v, want %v", got, tt.want)
			}
		})
	}
}
