package tests

import (
	"database/sql"
	"os"
	"testing"

	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/gkampitakis/go-snaps/snaps"
	"github.com/go-testfixtures/testfixtures/v3"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jbactad/loop/application"
	"github.com/jbactad/loop/application/ports"
	"github.com/jbactad/loop/graph/generated"
	"github.com/jbactad/loop/graph/resolvers"
	"github.com/jbactad/loop/infrastructure"
	"github.com/samber/do"

	_ "github.com/lib/pq"
)

func TestMain(t *testing.M) {
	err := setupDB()
	if err != nil {
		panic(err)
	}

	v := t.Run()

	// After all tests have run `go-snaps` can check for not used snapshots
	snaps.Clean(t)

	os.Exit(v)
}

func setupDB() (err error) {
	db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost:5432/loop?sslmode=disable")
	if err != nil {
		return
	}

	_, err = db.Exec("DROP SCHEMA public CASCADE; CREATE SCHEMA public;")
	if err != nil {
		return
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://../../db/migrations",
		"postgres",
		driver,
	)
	if err != nil {
		return
	}

	err = m.Up()
	if err != nil {
		return
	}

	fixtures, err := testfixtures.New(
		testfixtures.Database(db),
		testfixtures.Dialect("postgres"),
		testfixtures.Directory("../../db/seeds"),
		testfixtures.DangerousSkipTestDatabaseCheck(),
	)
	if err != nil {
		return err
	}

	err = fixtures.Load()
	return
}

func NewTestClient(t *testing.T, timeProvider ports.TimeProvider, uuidGenerator ports.UUIDGenerator) *client.Client {
	t.Helper()
	injector := do.NewWithOpts(&do.InjectorOpts{})

	err := infrastructure.ProvideDatabaseConnection(injector)
	if err != nil {
		t.Fatal(err)
	}

	do.Provide(injector, func(i *do.Injector) (ports.UUIDGenerator, error) {
		return uuidGenerator, nil
	})

	do.Provide(injector, func(i *do.Injector) (ports.TimeProvider, error) {
		return timeProvider, nil
	})

	infrastructure.ProvideRepositories(injector)
	application.ProvideQueryUseCases(injector)
	application.ProvideCommandUseCases(injector)
	resolvers.ProvideResolver(injector)
	if err != nil {
		t.Fatal(err)
	}

	r := do.MustInvoke[*resolvers.Resolver](injector)

	return client.New(handler.NewDefaultServer(generated.NewExecutableSchema(*generated.NewConfig(r))))
}
