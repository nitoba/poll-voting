package routes

import (
	"github.com/gin-gonic/gin"
	configs "github.com/nitoba/poll-voting/config"
	"github.com/nitoba/poll-voting/internal/infra/http/controllers"
)

func AuthRoutes(app *gin.Engine) {
	ctn := configs.GetContainer()
	registerController := ctn.Get("registerController").(*controllers.RegisterVoterController)

	router := app.Group("/auth")
	{
		router.POST("/register", registerController.Handle)
	}
}
