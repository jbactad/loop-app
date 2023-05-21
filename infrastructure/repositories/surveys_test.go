package repositories

import (
	"context"
	"reflect"
	"testing"

	"github.com/jbactad/loop/domain"
)

func TestSurveyRepository_GetSurveys(t *testing.T) {
	type fields struct {
		db Database
	}
	type args struct {
		ctx    context.Context
		limit  int
		offset int
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		wantSurveys []*domain.Survey
		wantErr     bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &SurveyRepository{
				db: tt.fields.db,
			}
			gotSurveys, err := repo.GetSurveys(tt.args.ctx, tt.args.limit, tt.args.offset)
			if (err != nil) != tt.wantErr {
				t.Errorf("SurveyRepository.GetSurveys() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotSurveys, tt.wantSurveys) {
				t.Errorf("SurveyRepository.GetSurveys() = %v, want %v", gotSurveys, tt.wantSurveys)
			}
		})
	}
}
