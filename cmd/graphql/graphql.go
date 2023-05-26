package graphql

import (
	"log"
	"net/http"
	"os"
	"syscall"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gookit/gcli/v3"
	"github.com/jbactad/loop/application"
	"github.com/jbactad/loop/graph/generated"
	"github.com/jbactad/loop/graph/resolvers"
	"github.com/jbactad/loop/infrastructure"
	"github.com/samber/do"
)

const defaultPort = "8081"

func NewRunServerCommand() *gcli.Command {
	return &gcli.Command{
		Name:    "graphql",
		Desc:    "Run GraphQL server.",
		Aliases: []string{"gql"},
		Func:    RunServer,
	}
}

func RunServer(cmd *gcli.Command, args []string) error {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	injector := do.NewWithOpts(&do.InjectorOpts{})

	err := ConfigureServices(injector)
	if err != nil {
		return err
	}

	r := do.MustInvoke[*resolvers.Resolver](injector)
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(*generated.NewConfig(r)))

	http.Handle("/", playground.ApolloSandboxHandler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
	injector.ShutdownOnSignals(os.Interrupt, os.Kill, syscall.SIGTERM)

	return nil
}

func ConfigureServices(injector *do.Injector) error {
	err := infrastructure.ProvideDatabaseConnection(injector)
	if err != nil {
		return err
	}

	infrastructure.ProvideRepositories(injector)
	application.ProvideQueryUseCases(injector)
	resolvers.ProvideResolver(injector)

	return nil
}
