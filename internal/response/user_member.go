package response

import "go-start/internal/pkg/helper"

type UserMemberShowRes struct {
	Id           string       `json:"id"`
	Passport     string       `json:"passport"`
	Password     string       `json:"-"`
	Email        string       `json:"email"`
	Status       string       `json:"status"`
	StatusText   string       `json:"status_text" gorm:"-"`
	IfVerify     string       `json:"if_verify"`
	IfVerifyText string       `json:"if_verify_text" gorm:"-"`
	Nickname     string       `json:"nickname"`
	CreatedAt    helper.FTime `json:"created_at"`
	UpdatedAt    helper.FTime `json:"updated_at"`
}
