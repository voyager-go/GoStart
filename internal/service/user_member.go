package service

import (
	"go-start/internal/model/entity"
	"go-start/internal/request"
	"go-start/internal/response"
)

// UserMemberManageService 后台玩家信息管理
type UserMemberManageService interface {
	List(req request.UserMemberListReq) *[]entity.UserMember
	ChangeStatus(req request.UserMemberChangeStatusReq) error
}

// UserMemberService 玩家信息交互
type UserMemberService interface {
	UpdateAvatar(id int64, avatar string) error
	Show(id int64) (res *response.UserMemberShowRes, err error)
	Suggest(req request.UserMemberSuggestReq) (res response.UserMemberSuggestListRes, err error)
	SignUp(req request.UserMemberSignUpReq) error
	SignIn(req request.UserMemberSignInReq) (token string, err error)
	GetNotVerify() []entity.UserMember
	VerifyEmail(req request.UserMemberVerifyEmailReq) error
}
