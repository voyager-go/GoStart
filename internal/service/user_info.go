package service

import (
	"go-start/internal/model/entity"
	"go-start/internal/request"
	"go-start/internal/response"
)

// UserInfoService 后台用户管理
type UserInfoService interface {
	Show(req request.UserInfoShowReq) (res *response.UserInfoShowRes, err error)
	List(req request.UserInfoListReq) *[]entity.UserInfo
	Create(req request.UserInfoCreateReq) error
}
