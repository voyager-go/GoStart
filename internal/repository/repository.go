package repository

import "go-start/internal/service"

type DataService struct {
	service.UserInfoService
}

func NewDataProvider() *DataService {
	return &DataService{
		UserInfoService: newUserInfoRepository(),
	}
}
