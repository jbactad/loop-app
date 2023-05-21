package infrastructure

import (
	"github.com/jbactad/loop/application/ports"
	"github.com/jbactad/loop/infrastructure/repositories"
	"github.com/samber/do"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ProvideRepositories(injector *do.Injector) {
	do.Provide(injector, func(i *do.Injector) (ports.SurveyProvider, error) {
		return repositories.NewSurveyRepository(do.MustInvoke[repositories.Database](i)), nil
	})
}

func ProvideDatabaseConnection(injector *do.Injector) error {
	// TODO: Move this to configuration
	dsn := "host=localhost user=postgres password=postgres dbname=loop port=5432 sslmode=disable TimeZone=UTC"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	do.Provide(injector, func(i *do.Injector) (repositories.Database, error) {
		return NewGormDB(db), nil
	})

	return nil
}
