package resolvers

import (
	"github.com/jbactad/loop/application/queries"
	"github.com/samber/do"
)

func ProvideResolver(injector *do.Injector) {
	do.Provide(injector, func(i *do.Injector) (*Resolver, error) {
		return &Resolver{
			Queries: do.MustInvoke[queries.UseCases](i),
		}, nil
	})
}
