package rest

import (
	"github.com/gin-contrib/cors"
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
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4321", "http://localhost:3333"},
		AllowMethods:     []string{"GET", "POST", "DELETE", "PUT", "PATCH", "OPTIONS"},
		AllowHeaders:     []string{"Content-Length", "Content-Type", "Accept-Encoding", "Authorization", "Cache-Control"},
		ExposeHeaders:    []string{"Origin"},
		AllowCredentials: true,
	}))

	// Register routes
	routes.AuthRoutes(r)
	routes.PollRoutes(r)

	// Register docs
	routes.DocsRoutes(r)

	return r
}
