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
	c := cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4321"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowCredentials: true,
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
	})
	r.Use(c)

	// Register routes
	routes.AuthRoutes(r)
	routes.PollRoutes(r)

	// Register docs
	routes.DocsRoutes(r)

	return r
}
