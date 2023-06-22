package generated_test

import (
	"os"
	"testing"

	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/gkampitakis/go-snaps/snaps"
	"github.com/jbactad/loop/application"
	"github.com/jbactad/loop/graph/generated"
	"github.com/jbactad/loop/graph/resolvers"
	"github.com/jbactad/loop/infrastructure"
	"github.com/samber/do"
)

func TestMain(t *testing.M) {
	v := t.Run()

	// After all tests have run `go-snaps` can check for not used snapshots
	snaps.Clean(t)

	os.Exit(v)
}

func NewTestClient(t *testing.T) *client.Client {
	t.Helper()
	injector := do.NewWithOpts(&do.InjectorOpts{})

	err := infrastructure.ProvideDatabaseConnection(injector)
	if err != nil {
		t.Fatal(err)
	}

	infrastructure.ProvideRepositories(injector)
	application.ProvideQueryUseCases(injector)
	resolvers.ProvideResolver(injector)
	if err != nil {
		t.Fatal(err)
	}

	r := do.MustInvoke[*resolvers.Resolver](injector)

	return client.New(handler.NewDefaultServer(generated.NewExecutableSchema(*generated.NewConfig(r))))
}
