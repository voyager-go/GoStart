package response

import "go-start/internal/pkg/helper"

type UserMemberShowRes struct {
	Id           int64        `json:"id"`
	Avatar       string       `json:"avatar"`
	Passport     string       `json:"passport"`
	Password     string       `json:"-"`
	Email        string       `json:"email"`
	Status       int          `json:"status"`
	StatusText   string       `json:"status_text" gorm:"-"`
	IfVerify     int          `json:"if_verify"`
	IfVerifyText string       `json:"if_verify_text" gorm:"-"`
	Nickname     string       `json:"nickname"`
	CreatedAt    helper.FTime `json:"created_at"`
	UpdatedAt    helper.FTime `json:"updated_at"`
}

type UserMemberSuggestListRes struct {
	List []UserMemberSuggest `json:"list"`
}
type UserMemberSuggest struct {
	Id       int64  `json:"id"`
	Passport string `json:"passport"`
	Nickname string `json:"nickname"`
}

type UserMember struct {
	Id       int64  `json:"id"`
	Avatar   string `json:"avatar"`
	Passport string `json:"passport"`
	Nickname string `json:"nickname"`
}
