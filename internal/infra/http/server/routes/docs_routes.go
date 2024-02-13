package routes

import (
	"github.com/gin-gonic/gin"
	_ "github.com/nitoba/poll-voting/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func DocsRoutes(app *gin.Engine) {
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
