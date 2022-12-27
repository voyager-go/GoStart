package routes

import (
	"github.com/gin-gonic/gin"
	"go-start/internal/controller"
)

func InitTopicRoutes(r *gin.RouterGroup) (router gin.IRoutes) {
	routes := r.Group("")
	{
		routes.GET("/topic/:topic_id", controller.Topic.Show)
	}
	return routes
}

func InitTopicPublicRoutes(r *gin.RouterGroup) (router gin.IRoutes) {
	routes := r.Group("")
	{
		routes.GET("/topic/list", controller.Topic.List)
	}
	return routes
}
