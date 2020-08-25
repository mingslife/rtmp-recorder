package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterAuthMiddleware(r *gin.RouterGroup, token string) {
	r.Use(func(ctx *gin.Context) {
		if ctx.DefaultQuery("token", "") == token {
			ctx.Next()
		} else {
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}
	})
}
