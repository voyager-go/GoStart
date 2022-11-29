package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	PublicMiddleware = []gin.HandlerFunc{
		cors.Default(),
		LoggerToFile(),
	}
)
