package generated_test

import (
	"encoding/json"
	"testing"

	"github.com/99designs/gqlgen/client"
)

func TestSurveys(t *testing.T) {
	c := NewTestClient(t)

	var resp struct {
		Surveys []struct {
			ID          string
			Description string
			Name        string
			Question    string
		}
	}

	c.MustPost(`query Surveys($limit: Int, $page: Int) {
  surveys(limit: $limit, page: $page) {
    id
    description
    name
    question
  }
}`,
		&resp,
		client.Var("limit", 10),
		client.Var("page", 0),
	)

	marshalled, err := json.MarshalIndent(resp.Surveys, "", "  ")
	if err != nil {
		t.Fatal(err)
	}

	Snapshoter.Snapshot(marshalled)
}
