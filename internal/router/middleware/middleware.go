package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func PublicMiddleware() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		cors.Default(),
		LoggerToFile(),
	}
}

func AllowPreCheck() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if strings.ToUpper(ctx.Request.Method) == "OPTIONS" {
			ctx.JSON(http.StatusOK, "ok!")
			ctx.Abort()
		}
	}
}
