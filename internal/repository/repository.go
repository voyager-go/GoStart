package repository

import "go-start/internal/service"

type DataService struct {
	service.UserInfoService
	service.UserMemberService
	service.UserMemberManageService
}

func NewDataProvider() *DataService {
	return &DataService{
		UserInfoService:         newUserInfoRepository(),
		UserMemberService:       newUserMemberRepository(),
		UserMemberManageService: newUserMemberManageRepository(),
	}
}
