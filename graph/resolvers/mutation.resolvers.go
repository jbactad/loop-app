package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.45

import (
	"context"
	"fmt"

	"github.com/jbactad/loop/application/commands"
	"github.com/jbactad/loop/graph/generated"
	"github.com/jbactad/loop/graph/models"
)

// CreateSurvey is the resolver for the createSurvey field.
func (r *mutationResolver) CreateSurvey(ctx context.Context, input models.NewSurvey) (*models.Survey, error) {
	s, err := r.Commands.CreateSurvey(ctx, commands.CreateSurveyCommand{
		Name:        input.Name,
		Description: input.Description,
		Question:    input.Question,
	})
	if err != nil {
		return nil, err
	}

	return SurveyToResponse(s), nil
}

// CreateSurveyResponse is the resolver for the createSurveyResponse field.
func (r *mutationResolver) CreateSurveyResponse(ctx context.Context, input models.NewSurveyResponse) (*models.SurveyResponse, error) {
	panic(fmt.Errorf("not implemented: CreateSurveyResponse - createSurveyResponse"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
