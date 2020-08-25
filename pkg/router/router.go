package router

import (
	"github.com/gin-gonic/gin"

	"rtmp-recorder/pkg/conf"
	"rtmp-recorder/pkg/controllers"
	"rtmp-recorder/pkg/middleware"
)

func NewRouter(c *conf.Config) *gin.Engine {
	router := gin.Default()

	apiRouter := router.Group("/api")
	middleware.RegisterCorsMiddleware(apiRouter)
	middleware.RegisterAuthMiddleware(apiRouter, c.Token)
	controllers.NewVideoController(apiRouter)

	return router
}
