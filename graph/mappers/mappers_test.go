package mappers

import (
	"reflect"
	"testing"

	"github.com/jbactad/loop/domain"
	"github.com/jbactad/loop/graph/models"
)

func TestSurveysToSurveysResponse(t *testing.T) {
	type args struct {
		surveys []*domain.Survey
	}
	tests := []struct {
		name string
		args args
		want []*models.Survey
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SurveysToSurveysResponse(tt.args.surveys); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SurveysToSurveysResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSurveyToResponse(t *testing.T) {
	type args struct {
		survey *domain.Survey
	}
	tests := []struct {
		name string
		args args
		want *models.Survey
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SurveyToResponse(tt.args.survey); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SurveyToResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}
