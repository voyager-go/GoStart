package routes

import (
	"github.com/gin-gonic/gin"
	"go-start/internal/controller"
)

func InitUserMemberRoutes(r *gin.RouterGroup) (router gin.IRoutes) {
	userRoutes := r.Group("")
	{
		userRoutes.GET("/user-member/:user_member_id", controller.UserMember.Show)
		userRoutes.POST("/user-member/sign-logout", controller.UserMember.SignIn)
	}
	return userRoutes
}
