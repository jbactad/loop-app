package repositories

import (
	"time"

	"github.com/jbactad/loop/domain"
)

type SurveyResponsesData struct {
	ID        string      `gorm:"primaryKey"`
	Survey    *SurveyData `gorm:"foreignKey:SurveyID"`
	Answer    string
	Rating    int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (sr SurveyResponsesData) ToDomain() *domain.SurveyResponse {
	return domain.NewSurveyResponse(sr.ID, sr.Survey.ToDomain(), sr.Answer, sr.Rating, sr.CreatedAt, sr.UpdatedAt)
}
