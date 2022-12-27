package crontab

import (
	"github.com/robfig/cron/v3"
)

func Start() {
	c := cron.New(cron.WithParser(cron.NewParser(cron.SecondOptional | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor)))
	c.AddFunc("@every 10s", PushNewUserMemberToTaskQueue)
	c.Start()
}
