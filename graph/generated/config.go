package generated

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/jbactad/loop/graph/models"
)

func NewConfig(resolver ResolverRoot) *Config {
	c := &Config{
		Resolvers: resolver,
	}

	c.Directives.HasRole = func(ctx context.Context, obj interface{}, next graphql.Resolver, role models.Role) (res interface{}, err error) {
		return next(ctx)
	}

	return c
}
