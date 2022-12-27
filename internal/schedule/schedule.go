package main

import (
	"github.com/hibiken/asynq"
	"go-start/internal/consts"
	"go-start/internal/schedule/task"
	"log"
)

func main() {
	redisAddr := "127.0.0.1:6379"
	srv := asynq.NewServer(asynq.RedisClientOpt{Addr: redisAddr}, asynq.Config{
		Concurrency: 10,
		Queues: map[string]int{
			consts.AsynqCritical: 6,
			consts.AsynqDefault:  3,
			consts.AsynqLow:      1,
		},
	})

	mux := asynq.NewServeMux()
	mux.HandleFunc(task.TypeEmailDelivery, task.HandleEmailDeliveryTask)
	if err := srv.Run(mux); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}
