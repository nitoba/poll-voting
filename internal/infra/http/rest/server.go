package rest

import (
	"github.com/gin-gonic/gin"
	configs "github.com/nitoba/poll-voting/config"
	"github.com/nitoba/poll-voting/internal/infra/http/rest/routes"
)

func GetServer() *gin.Engine {
	conf := configs.GetConfig()

	println("Running in " + conf.ENV)

	if conf.ENV == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	if conf.ENV == "test" {
		gin.SetMode(gin.TestMode)
	}

	r := gin.Default()

	// Register routes
	routes.AuthRoutes(r)
	routes.PollRoutes(r)

	// Register docs
	routes.DocsRoutes(r)

	return r
}
