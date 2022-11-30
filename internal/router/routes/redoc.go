package routes

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	docs "go-start/assets/docs"
	"go-start/config"
	"go-start/internal/pkg/redoc"
)

func InitDocRoutes(r *gin.RouterGroup) (router gin.IRoutes) {
	docRoutes := r.Group("")
	{
		docRoutes.GET("/redoc", func(ctx *gin.Context) {
			doc := &redoc.Redoc{
				Title:    config.Cfg.Redoc.Title,
				DocPath:  config.Cfg.Redoc.DocPath,
				SpecPath: config.Cfg.Redoc.SpecPath,
				SpecFile: config.Cfg.Redoc.SpecFile,
				Desc:     config.Cfg.Redoc.Desc,
			}
			h := doc.Handler()
			h(ctx.Writer, ctx.Request)
		})
		docRoutes.StaticFile("/swagger.json", "assets/docs/swagger.json")
		docs.SwaggerInfo.BasePath = "/api"
		docRoutes.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
	return docRoutes
}
