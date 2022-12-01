package routes

import (
	"github.com/gin-gonic/gin"
	"go-start/internal/controller"
)

func InitUserRoutes(r *gin.RouterGroup) (router gin.IRoutes) {
	userRoutes := r.Group("")
	{
		userRoutes.GET("/user/:id", controller.User.Show)
		userRoutes.POST("/user", controller.User.Create)
	}
	return userRoutes
}
