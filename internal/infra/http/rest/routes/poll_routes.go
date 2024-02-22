package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nitoba/poll-voting/internal/infra/cryptography"
	"github.com/nitoba/poll-voting/internal/infra/database/prisma/repositories"
	"github.com/nitoba/poll-voting/internal/infra/http/controllers"
	"github.com/nitoba/poll-voting/internal/infra/http/rest/middlewares"
	"github.com/nitoba/poll-voting/pkg/di"
)

func PollRoutes(app *gin.Engine) {
	ctn := di.GetContainer()
	jwtEncrypter := ctn.Get("encrypter").(*cryptography.JWTEncrypter)
	voterRepository := ctn.Get("voterRepository").(*repositories.VotersRepositoryPrisma)
	createPollController := ctn.Get("createPollController").(*controllers.CreatePollController)
	fetchPollsController := ctn.Get("fetchPollsController").(*controllers.FetchPollsController)
	fetchPollsByOwnerController := ctn.Get("fetchPollsByOwnerController").(*controllers.FetchPollsByOwnerController)
	getPollByIdController := ctn.Get("getPollByIdController").(*controllers.GetPollByIdController)
	voteOnPollController := ctn.Get("voteOnPollController").(*controllers.VoteOnPollController)
	updateCountingVotesController := ctn.Get("updateCountingVotesController").(*controllers.UpdateCountingVotesController)

	router := app.Group("/polls").Use(middlewares.AuthRequiredCookie(jwtEncrypter, voterRepository))
	{
		router.GET("/all", fetchPollsController.Handle)
		router.GET("/", fetchPollsByOwnerController.Handle)
		router.GET("/:id", getPollByIdController.Handle)
		router.POST("/", createPollController.Handle)
		router.POST("/:id/vote", voteOnPollController.Handle)
	}

	app.GET("polls/:id/results", updateCountingVotesController.Handle)
}
