package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nitoba/poll-voting/internal/infra/database"
	"github.com/nitoba/poll-voting/internal/infra/http"
	"github.com/nitoba/poll-voting/internal/infra/http/controllers"
	"github.com/sarulabs/di"
)

func AuthRoutes(app *gin.Engine) {
	databaseModule := database.NewDatabaseModule()
	httpModule := http.NewHttpModule()
	builder, _ := di.NewBuilder()
	builder.Add(databaseModule.PrismaDB)
	builder.Add(httpModule.Hasher, httpModule.VoterRepository, httpModule.RegisterVoterUseCase, httpModule.RegisterController)
	ctn := builder.Build()
	registerController := ctn.Get("registerController").(*controllers.RegisterVoterController)

	router := app.Group("/auth")
	{
		router.POST("/register", registerController.Handle)
	}
}
