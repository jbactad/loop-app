package resolvers

import (
	"context"
	"reflect"
	"testing"

	"github.com/jbactad/loop/graph/models"
)

func Test_mutationResolver_CreateSurvey(t *testing.T) {
	type fields struct {
		Resolver *Resolver
	}
	type args struct {
		ctx   context.Context
		input models.NewSurvey
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.Survey
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &mutationResolver{
				Resolver: tt.fields.Resolver,
			}
			got, err := r.CreateSurvey(tt.args.ctx, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("mutationResolver.CreateSurvey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mutationResolver.CreateSurvey() = %v, want %v", got, tt.want)
			}
		})
	}
}
