package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"go-start/config"
)

var Client *redis.Client

func NewRedis() {
	cfg := config.Cfg.Redis
	Client = redis.NewClient(&redis.Options{
		Addr:     cfg.Host,
		Password: cfg.Password,
		DB:       cfg.DbNum,
	})
	if err := Client.Ping(context.Background()).Err(); err != nil {
		panic(fmt.Errorf("无法连接Redis，请先检查Redis配置信息，错误详情为: %s", err.Error()))
	}
}
