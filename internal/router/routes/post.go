package routes

import (
	"github.com/gin-gonic/gin"
	"go-start/internal/controller"
)

func InitPostRoutes(r *gin.RouterGroup) (router gin.IRoutes) {
	routes := r.Group("/post")
	{
		routes.GET("/show", controller.Post.Show)
		routes.GET("/comment-list", controller.Post.CommentList)
		routes.POST("/comment-create", controller.Post.CommentCreate)
		routes.POST("/publish", controller.Post.Publish)
		routes.POST("/star", controller.Post.Star)
		routes.POST("/collect", controller.Post.Collect)
	}
	return routes
}

func InitPostPublicRoutes(r *gin.RouterGroup) (router gin.IRoutes) {
	routes := r.Group("")
	{
		routes.GET("/post/list", controller.Post.List)
	}
	return routes
}
