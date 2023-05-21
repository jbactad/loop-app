package application

import (
	"github.com/jbactad/loop/application/ports"
	"github.com/jbactad/loop/application/queries"
	"github.com/mehdihadeli/go-mediatr"
	"github.com/samber/do"
)

func ProvideQueryHandlers(injector *do.Injector) {
	do.Provide(injector, func(i *do.Injector) (*queries.GetSurveysQueryHandler, error) {
		sp := do.MustInvoke[ports.SurveyProvider](i)

		return queries.NewGetSurveysQueryHandler(sp), nil
	})

	mediatr.RegisterRequestHandler[queries.GetSurveysQuery, queries.GetSurveysQueryResponse](do.MustInvoke[*queries.GetSurveysQueryHandler](injector))
}
