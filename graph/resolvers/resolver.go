package resolvers

import (
	"github.com/jbactad/loop/application/commands"
	"github.com/jbactad/loop/application/queries"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Commands commands.UseCases
	Queries  queries.UseCases
}
