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
	createPollController := ctn.Get("createPollController").(*controllers.CreatePollController)
	jwtEncrypter := ctn.Get("encrypter").(*cryptography.JWTEncrypter)
	voterRepository := ctn.Get("voterRepository").(*repositories.VotersRepositoryPrisma)

	router := app.Group("/polls")
	router.Use(middlewares.AuthRequired(jwtEncrypter, voterRepository))
	{
		router.POST("/", createPollController.Handle)
	}
}
