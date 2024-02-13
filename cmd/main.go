package main

import (
	configs "github.com/nitoba/poll-voting/config"
	"github.com/nitoba/poll-voting/internal/infra/database"
	"github.com/nitoba/poll-voting/internal/infra/http"
)

func main() {
	configs.InitContainer()

	databaseModule := database.NewDatabaseModule()
	httpModule := http.NewHttpModule()

	configs.RegisterDependency(databaseModule.GetDependencies()...)
	configs.RegisterDependency(httpModule.GetDependencies()...)

	configs.BuildDependencies()
}
