package asynq_client

import (
	"github.com/hibiken/asynq"
	"go-start/config"
)

var Asynq *asynq.Client

func NewAsynq() {
	redisAddr := config.Cfg.Redis.Host + ":" + config.Cfg.Redis.Port
	Asynq = asynq.NewClient(asynq.RedisClientOpt{Addr: redisAddr})
}
