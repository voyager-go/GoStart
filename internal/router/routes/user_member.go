package routes

import (
	"github.com/gin-gonic/gin"
	"go-start/internal/controller"
)

func InitUserMemberRoutes(r *gin.RouterGroup) (router gin.IRoutes) {
	userRoutes := r.Group("")
	{
		userRoutes.GET("/user-member/:id", controller.User.Show)
		userRoutes.POST("/user-member/sign-in", controller.User.Create)
		userRoutes.POST("/user-member/sign-up", controller.UserMember.SignUp)
		userRoutes.POST("/user-member/sign-logout", controller.User.Create)
	}
	return userRoutes
}
