package repository

import "go-start/internal/service"

type DataService struct {
	service.UserInfoService
	service.UserMemberService
	service.UserMemberManageService
	service.TopicService
	service.PostService
	service.PostManageService
}

func NewDataProvider() *DataService {
	return &DataService{
		UserInfoService:         newUserInfoRepository(),
		UserMemberService:       newUserMemberRepository(),
		UserMemberManageService: newUserMemberManageRepository(),
		TopicService:            newTopicRepository(),
		PostService:             newPostRepository(),
		PostManageService:       newPostManageRepository(),
	}
}
