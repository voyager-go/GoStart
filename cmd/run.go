package cmd

import (
	"GoStart/config"
	"GoStart/internal/pkg/log"
	"GoStart/internal/pkg/response"
	"GoStart/internal/pkg/validator_trans"
	"GoStart/internal/router/middleware"
	"GoStart/internal/router/routes"
	"github.com/gin-gonic/gin"
	"github.com/urfave/cli/v2"
)

var App = &cli.App{
	Name:     "main",
	Usage:    "start this project",
	Commands: []*cli.Command{},
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:        "env",
			Value:       "dev",
			Usage:       "请选择配置文件 [dev | pre]",
			Destination: &config.ConfEnv,
		},
		&cli.StringFlag{
			Name:        "port",
			Value:       config.AppPort,
			Usage:       "请选择启动端口",
			Destination: &config.AppPort,
		},
	},
	Before: func(*cli.Context) error {
		// 初始化配置文件
		config.NewConfig()
		// 初始化日志追踪
		log.NewLogger()
		// 初始化验证器翻译
		validator_trans.NewTrans()
		return nil
	},
	Action: func(*cli.Context) error {
		var (
			srv = gin.New()
			r   = new(response.R)
		)
		// 404 处理
		srv.NoRoute(func(ctx *gin.Context) {
			r.RequestNotFound(ctx)
		})
		// 路由分组
		normalGroup := srv.Group("/api", middleware.PublicMiddleware...)
		// 用户组
		routes.InitUserRoutes(normalGroup)
		// 启动项目
		return srv.Run(":" + config.AppPort)
	},
}
