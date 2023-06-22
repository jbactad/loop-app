package repositories

import (
	"context"
	"time"

	"github.com/jbactad/loop/domain"
)

type SurveyData struct {
	ID          string `gorm:"primaryKey"`
	Name        string
	Description string
	Question    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (s SurveyData) ToDomain() *domain.Survey {
	return domain.NewSurvey(s.ID, s.Name, s.Description, s.Question, s.CreatedAt, s.UpdatedAt)
}

type SurveyRepository struct {
	db Database
}

func NewSurveyRepository(dbConn Database) *SurveyRepository {
	return &SurveyRepository{
		db: dbConn,
	}
}

func (repo *SurveyRepository) GetSurveys(
	ctx context.Context, limit int, offset int,
) (surveys []*domain.Survey, err error) {
	var surveyDatas []*SurveyData
	err = repo.db.
		Table("surveys").
		Limit(limit).
		Offset(offset).
		Find(&surveyDatas).
		Error()
	if err != nil {
		return nil, err
	}

	for _, surveyData := range surveyDatas {
		surveys = append(surveys, surveyData.ToDomain())
	}

	return surveys, nil
}

func (f *SurveyRepository) CreateSurvey(ctx context.Context, survey *domain.Survey) error {
	panic("not implemented") // TODO: Implement
}
