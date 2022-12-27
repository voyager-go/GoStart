package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/urfave/cli/v2"
	"go-start/config"
	"go-start/internal/controller"
	"go-start/internal/crontab"
	"go-start/internal/pkg/asynq_client"
	"go-start/internal/pkg/log"
	"go-start/internal/pkg/mysql"
	"go-start/internal/pkg/redis"
	"go-start/internal/pkg/res"
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
		// 初始化数据库连接
		mysql.NewMysql()
		// 初始化Redis连接
		redis.NewRedis()
		// 初始化验证器翻译
		validator_trans.NewTrans()
		// 初始化异步任务队列
		asynq_client.NewAsynq()
		return nil
	},
	Action: func(*cli.Context) error {
		var (
			srv = gin.New()
			pm  = middleware.PublicMiddleware()
			ap  = middleware.AllowPreCheck()
			r   = new(res.R)
		)
		srv.Use(ap)
		// 404 处理
		srv.NoRoute(func(ctx *gin.Context) {
			r.RequestNotFound(ctx)
		})
		// 路由分组 - 公共路由
		normalGroup := srv.Group("/api", pm...)
		// 路由分组 - 授权路由
		authGroup := srv.Group("/api", append(pm, middleware.Auth)...)
		// 初始化控制层基础方法
		controller.InitMethods()
		// 用户组
		routes.InitUserInfoRoutes(normalGroup)
		// 文档组
		routes.InitDocRoutes(normalGroup)
		// 认证组
		routes.InitAuthRoutes(normalGroup)
		// 玩家组
		routes.InitUserMemberRoutes(authGroup)
		// 玩家开放组
		routes.InitUserMemberPublicRoutes(normalGroup)
		// 主题开放组
		routes.InitTopicPublicRoutes(normalGroup)
		// 主题认证组
		routes.InitTopicRoutes(authGroup)
		// 推文开放组
		routes.InitPostPublicRoutes(normalGroup)
		// 静态文件组
		routes.InitStaticRoutes(normalGroup)
		// 推文认证组
		routes.InitPostRoutes(authGroup)
		// 开启定时任务
		crontab.Start()
		// 生成swagger文档，生成失败时无法捕捉日志，建议手动执行[swag init --output assets/docs]
		if config.Cfg.Cmd.SwagName != "" {
			cmd := exec.Command(config.Cfg.Cmd.SwagName, config.Cfg.Cmd.SwagArgs...)
			_ = cmd.Run()
		}
		// 启动项目
		return srv.Run(":" + config.AppPort)
	},
}
