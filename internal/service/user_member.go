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
	Show(req request.UserMemberShowReq) (res *response.UserMemberShowRes, err error)
	SignUp(req request.UserMemberSignUpReq) error
	SignIn(req request.UserMemberSignInReq) (token string, err error)
}
