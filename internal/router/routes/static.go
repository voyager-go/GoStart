package routes

import (
	"github.com/gin-gonic/gin"
	"go-start/config"
	"net/http"
)

func InitStaticRoutes(r *gin.RouterGroup) (router gin.IRoutes) {
	routes := r.Group("")
	{
		routes.StaticFS("/static", http.Dir(config.Cfg.Server.UploadPath))
	}
	return routes
}
