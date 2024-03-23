package tests_test

import (
	"testing"

	"github.com/99designs/gqlgen/client"
	"github.com/gkampitakis/go-snaps/snaps"
	"github.com/jbactad/loop/application/ports/mocks"
	"github.com/stretchr/testify/assert"
)

func TestSurveys(t *testing.T) {
	type args struct {
		query string
		limit int
		page  int
	}
	tests := []struct {
		name    string
		args    args
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "with valid query, it should return surveys without error",
			args: args{
				query: `query Surveys($limit: Int, $page: Int) {
  surveys(limit: $limit, page: $page) {
    id
    description
    name
    question
  }
}`,
				limit: 1,
				page:  0,
			},
		},
		{
			name: "with invalid query, it should return error",
			args: args{
				query: `query Surveys($limit: Int, $page: Int) {
					  surveys(limit: $limit, page: $page) {
						id
						description
						name
						question

						# This is invalid
						invalid
					  }
					}`,
				limit: 10,
				page:  0,
			},
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewTestClient(t, mocks.NewTimeProvider(t), mocks.NewUUIDGenerator(t))

			var resp struct {
				Surveys []struct {
					ID          string
					Description string
					Name        string
					Question    string
					CreatedAt   string
					UpdatedAt   string
				}
			}

			err := c.Post(tt.args.query,
				&resp,
				client.Var("limit", tt.args.limit),
				client.Var("page", tt.args.page),
			)
			if tt.wantErr != nil {
				tt.wantErr(t, err)
				return
			}

			snaps.MatchJSON(t, resp)
		})
	}
}
