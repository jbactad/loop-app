package db

import (
	"database/sql"

	"github.com/golang-migrate/migrate/v4"
	"github.com/gookit/gcli/v3"

	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/golang-migrate/migrate/v4/source/github"

	"github.com/go-testfixtures/testfixtures/v3"
	_ "github.com/lib/pq"
)

func NewDbCommand() *gcli.Command {
	return &gcli.Command{
		Name:    "db",
		Desc:    "Useful database commands.",
		Aliases: []string{"d"},
		Subs: []*gcli.Command{
			{
				Name:    "migrate",
				Desc:    "Execute migration scripts.",
				Aliases: []string{"m"},
				Func:    Migrate,
			},
			{
				Name:    "rollback",
				Desc:    "Rollback migration scripts.",
				Aliases: []string{"r"},
				Func:    Rollback,
			},
			{
				Name:    "seed",
				Desc:    "Load test data into the database.",
				Aliases: []string{"s"},
				Func:    Seed,
			},
		},
		Func: func(cmd *gcli.Command, args []string) error { return nil },
	}
}

func Migrate(cmd *gcli.Command, args []string) error {
	// TODO: Move this to configuration
	db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost:5432/loop?sslmode=disable")
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file://db/migrations",
		"postgres",
		driver,
	)
	if err != nil {
		return err
	}

	err = m.Up()

	return err
}

func Rollback(cmd *gcli.Command, args []string) error {
	// TODO: Move this to configuration
	db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost:5432/loop?sslmode=disable")
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file://db/migrations",
		"postgres",
		driver,
	)
	if err != nil {
		return err
	}

	err = m.Down()

	return err
}

func Seed(cmd *gcli.Command, args []string) error {
	// TODO: Move this to configuration
	db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost:5432/loop?sslmode=disable")
	if err != nil {
		return err
	}

	fixtures, err := testfixtures.New(
		testfixtures.Database(db),
		testfixtures.Dialect("postgres"),
		testfixtures.Directory("db/seeds"),
		testfixtures.DangerousSkipTestDatabaseCheck(),
	)
	if err != nil {
		return err
	}

	return fixtures.Load()
}
