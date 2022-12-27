package task

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hibiken/asynq"
	"go-start/internal/pkg/helper"
	"go-start/internal/pkg/log"
)

const (
	TypeEmailDelivery = "email:deliver"
)

type EmailDeliveryPayload struct {
	Email        string
	Subject      string
	TemplatePath string
	Link         string
}

func NewEmailDeliveryTask(userEmailInfo EmailDeliveryPayload) (*asynq.Task, error) {
	payload, err := json.Marshal(userEmailInfo)
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeEmailDelivery, payload), nil
}

func HandleEmailDeliveryTask(ctx context.Context, t *asynq.Task) error {
	var p EmailDeliveryPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		log.Logger.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
		return err
	}
	fmt.Println("email info: ", p.Email, p.Subject, p.TemplatePath, p.Link)
	err := helper.SendEmail(p.Email, p.Subject, p.TemplatePath, p.Link)
	if err != nil {
		log.Logger.Errorf("helper.SendEmail failed: %v: %w", err, asynq.SkipRetry)
		return err
	}

	return nil
}
