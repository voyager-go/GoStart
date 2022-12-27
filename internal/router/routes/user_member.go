package routes

import (
	"github.com/gin-gonic/gin"
	"go-start/internal/controller"
)

func InitUserMemberRoutes(r *gin.RouterGroup) (router gin.IRoutes) {
	userRoutes := r.Group("/user-member")
	{
		userRoutes.GET("/show", controller.UserMember.Show)
		userRoutes.GET("/suggest", controller.UserMember.Suggest)
		userRoutes.POST("/sign-logout", controller.UserMember.SignIn)
		userRoutes.POST("/upload-avatar", controller.UserMember.UploadAvatar)
	}
	return userRoutes
}

func InitUserMemberPublicRoutes(r *gin.RouterGroup) (router gin.IRoutes) {
	userRoutes := r.Group("/user-member")
	{
		userRoutes.GET("/verify-email", controller.UserMember.VerifyEmail)
	}
	return userRoutes
}
