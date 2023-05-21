package resolvers

import (
	"context"
	"reflect"
	"testing"

	"github.com/jbactad/loop/graph/models"
)

func Test_queryResolver_Surveys(t *testing.T) {
	type fields struct {
		Resolver *Resolver
	}
	type args struct {
		ctx   context.Context
		limit *int
		page  *int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*models.Survey
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &queryResolver{
				Resolver: tt.fields.Resolver,
			}
			got, err := r.Surveys(tt.args.ctx, tt.args.limit, tt.args.page)
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
