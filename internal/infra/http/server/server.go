package server

import (
	"github.com/gin-gonic/gin"
	"github.com/nitoba/poll-voting/internal/infra/http/server/routes"
)

func GetServer() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	// Register routes
	routes.AuthRoutes(r)

	return r
}
