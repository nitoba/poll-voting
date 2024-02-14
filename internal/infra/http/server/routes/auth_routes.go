package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nitoba/poll-voting/internal/infra/http/controllers"
	"github.com/nitoba/poll-voting/pkg/di"
)

func AuthRoutes(app *gin.Engine) {
	ctn := di.GetContainer()
	registerController := ctn.Get("registerController").(*controllers.RegisterVoterController)

	router := app.Group("/auth")
	{
		router.POST("/register", registerController.Handle)
	}
}
