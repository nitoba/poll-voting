package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nitoba/poll-voting/internal/infra/cryptography"
	"github.com/nitoba/poll-voting/internal/infra/database/prisma/repositories"
	"github.com/nitoba/poll-voting/internal/infra/http/controllers"
	"github.com/nitoba/poll-voting/internal/infra/http/server/middlewares"
	"github.com/nitoba/poll-voting/pkg/di"
)

func PollRoutes(app *gin.Engine) {
	ctn := di.GetContainer()
	jwtEncrypter := ctn.Get("encrypter").(*cryptography.JWTEncrypter)
	voterRepository := ctn.Get("voterRepository").(*repositories.VotersRepositoryPrisma)
	createPollController := ctn.Get("createPollController").(*controllers.CreatePollController)
	fetchPollsController := ctn.Get("fetchPollsController").(*controllers.FetchPollsController)

	router := app.Group("/polls")
	router.Use(middlewares.AuthRequired(jwtEncrypter, voterRepository))
	{
		router.GET("/", fetchPollsController.Handle)
		router.POST("/", createPollController.Handle)
	}
}
