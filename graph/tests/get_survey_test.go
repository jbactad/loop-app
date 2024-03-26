package tests

import (
	"testing"

	"github.com/99designs/gqlgen/client"
	"github.com/gkampitakis/go-snaps/snaps"
	"github.com/jbactad/loop/application/ports/mocks"
	"github.com/stretchr/testify/assert"
)

func TestSurvey(t *testing.T) {
	type args struct {
		query string
		id    string
	}
	testCases := []struct {
		name    string
		args    args
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "given valid query, should return surveys without error",
			args: args{
				query: `query Survey($surveyId: ID!) {
					survey(id: $surveyId) {
						id
						name
						description
						question
						createdAt
						updatedAt
					}
				}`,
				id: "8138abe3-06b1-5882-88dd-ba931b388f1d",
			},
			wantErr: assert.NoError,
		},
		{
			name: "given invalid query, should return error",
			args: args{
				query: `query Survey($surveyId: ID!) {
					survey(id: $surveyId) {
						id
						name
						description
						question
						createdAt
						updatedAt
					}
				}`,
				id: "",
			},
			wantErr: func(tt assert.TestingT, err error, i ...interface{}) bool {
				assert.Error(tt, err)
				assert.Contains(tt, err.Error(), "graphql: missing required query variable")
				return true
			},
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			c := NewTestClient(t, mocks.NewTimeProvider(t), mocks.NewUUIDGenerator(t))

			var resp struct {
				Survey struct {
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
				client.Var("surveyId", tt.args.id),
			)
			if !tt.wantErr(t, err) {
				return
			}

			snaps.MatchJSON(t, resp)
		})
	}
}
