package main

import (
	"github.com/gookit/gcli/v3"
	"github.com/jbactad/loop/cmd/db"
	"github.com/jbactad/loop/cmd/graphql"
)

// for test run: go build ./_examples/cliapp.go && ./cliapp
func main() {
	app := gcli.NewApp()
	app.Version = "1.0.0"
	app.Desc = "A simple pulse-surveys application."

	app.Add(graphql.NewRunServerCommand())
	app.Add(db.NewDbCommand())

	app.Run(nil)
}
