package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func PublicMiddleware() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		cors.Default(),
		LoggerToFile(),
	}
}
