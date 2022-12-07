package routes

import (
	"github.com/gin-gonic/gin"
	"go-start/internal/controller"
)

func InitAuthRoutes(r *gin.RouterGroup) (router gin.IRoutes) {
	userRoutes := r.Group("/auth")
	{
		userRoutes.POST("/sign-up", controller.UserMember.SignUp)
		userRoutes.POST("/sign-in", controller.UserMember.SignIn)
	}
	return userRoutes
}
