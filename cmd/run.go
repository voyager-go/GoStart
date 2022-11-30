package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/urfave/cli/v2"
	"go-start/config"
	"go-start/internal/pkg/log"
	"go-start/internal/pkg/response"
	"go-start/internal/pkg/validator_trans"
	"go-start/internal/router/middleware"
	"go-start/internal/router/routes"
	"os/exec"
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
			pm  = middleware.PublicMiddleware()
			r   = new(response.R)
		)
		// 404 处理
		srv.NoRoute(func(ctx *gin.Context) {
			r.RequestNotFound(ctx)
		})
		// 路由分组
		normalGroup := srv.Group("/api", pm...)
		// 用户组
		routes.InitUserRoutes(normalGroup)
		routes.InitDocRoutes(normalGroup)
		// 生成swagger文档
		cmd := exec.Command(config.Cfg.Cmd.SwagName, config.Cfg.Cmd.SwagArgs...)
		_ = cmd.Run()
		// 启动项目
		return srv.Run(":" + config.AppPort)
	},
}
