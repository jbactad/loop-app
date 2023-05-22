package application

import (
	"github.com/jbactad/loop/application/ports"
	"github.com/jbactad/loop/application/queries"
	"github.com/samber/do"
)

func ProvideQueryUseCases(injector *do.Injector) {
	do.Provide(injector, func(i *do.Injector) (queries.UseCases, error) {
		sp := do.MustInvoke[ports.SurveyProvider](i)

		return queries.New(sp), nil
	})
}
