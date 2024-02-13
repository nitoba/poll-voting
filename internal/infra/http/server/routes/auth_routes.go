package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nitoba/poll-voting/internal/infra/http/factories"
)

func AuthRoutes(app *gin.Engine) {
	registerController := factories.MakeRegisterController()

	router := app.Group("/auth")
	{
		router.POST("/register", registerController.Handle)
	}
}
