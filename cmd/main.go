package main

import (
	configs "github.com/nitoba/poll-voting/config"
	"github.com/nitoba/poll-voting/internal/infra"
	"github.com/nitoba/poll-voting/internal/infra/database/prisma"
	"github.com/nitoba/poll-voting/internal/infra/http/server"
	"github.com/nitoba/poll-voting/pkg/di"
)

// @title           Poll Voting API
// @version         1.0
// @description     A Poll Voting API in Golang
// @termsOfService  https://swagger.io/terms/

// @contact.name   Bruno Alves
// @contact.url    https://github.com/nitoba
// @contact.email  nito.ba.dev@gmail.com

// @license.name   NitoDev
// @license.url    https://github.com/nitoba

// @host      localhost:3333
// @BasePath  /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	configs.LoadConfig()

	prisma.Connect()
	defer prisma.Disconnect()

	di.InitContainer()

	app := infra.NewAppModule()

	di.RegisterModuleProviders(app.Providers)

	di.BuildDependencies()

	server := server.GetServer()

	server.Run(":" + configs.GetConfig().Port)
}
