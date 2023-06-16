package commands

import (
	"context"
	"reflect"
	"testing"

	"github.com/jbactad/loop/domain"
)

func TestCreateSurvey(t *testing.T) {
	type args struct {
		ctx context.Context
		cmd CreateSurveyCommand
	}
	tests := []struct {
		name    string
		args    args
		want    domain.Survey
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateSurvey(tt.args.ctx, tt.args.cmd)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateSurvey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateSurvey() = %v, want %v", got, tt.want)
			}
		})
	}
}
