package crontab

import (
	"errors"
	"fmt"
	"github.com/hibiken/asynq"
	"go-start/config"
	"go-start/internal/consts"
	"go-start/internal/pkg/asynq_client"
	"go-start/internal/pkg/log"
	"go-start/internal/repository"
	"go-start/internal/schedule/task"
	"time"
)

func PushNewUserMemberToTaskQueue() {
	users := repository.NewDataProvider().UserMemberService.GetNotVerify()
	for _, user := range users {
		t, err := task.NewEmailDeliveryTask(task.EmailDeliveryPayload{
			Email:        user.Email,
			Subject:      "欢迎注册推书小站",
			TemplatePath: config.Cfg.Server.AssetsPath + "email/tpl.html",
			Link:         fmt.Sprintf(consts.UserMemberVerifyLink, user.Id, user.VerifyCode),
		})
		if err != nil {
			log.Logger.Errorf("could not create task: %v", err)
			continue
		}
		q, err := asynq_client.Asynq.Enqueue(t, asynq.Queue(consts.AsynqCritical), asynq.Unique(24*time.Hour))
		if err != nil {
			if errors.Is(err, asynq.ErrDuplicateTask) {
				log.Logger.Errorf("the task already exist: %v", err)
			} else {
				log.Logger.Errorf("could not enqueue task: %v", err)
			}
			continue
		}
		log.Logger.Info(q.ID, q.Queue, q.State, string(q.Payload), q.CompletedAt, q.Group, q.Deadline, q.MaxRetry, q.Result)
	}
}
