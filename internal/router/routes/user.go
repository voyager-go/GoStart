package routes

import (
	"GoStart/internal/controller"
	"github.com/gin-gonic/gin"
)

func InitUserRoutes(r *gin.RouterGroup) (router gin.IRoutes) {
	userRoutes := r.Group("")
	{
		userRoutes.GET("/user/:uuid", controller.User.Show)
		userRoutes.POST("/user", controller.User.Create)
	}
	return userRoutes
}
