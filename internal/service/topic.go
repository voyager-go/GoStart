package service

import (
	"go-start/internal/request"
	"go-start/internal/response"
)

type TopicManageService interface {
}

type TopicService interface {
	Show(req request.TopicShowReq) (topic *response.TopicShowRes, err error)
	List(req request.TopicListReq) (list response.TopicListRes)
}
