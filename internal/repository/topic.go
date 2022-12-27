package repository

import (
	"errors"
	"go-start/internal/consts"
	"go-start/internal/model/entity"
	"go-start/internal/model/enum"
	"go-start/internal/pkg/helper"
	"go-start/internal/pkg/mysql"
	"go-start/internal/request"
	"go-start/internal/response"
	"go-start/internal/service"
	"gorm.io/gorm"
)

var (
	_ service.TopicService = (*topicRepository)(nil)
)

type topicRepository struct {
	db *gorm.DB
}

func newTopicRepository() service.TopicService {
	return &topicRepository{db: mysql.Conn}
}

func (r *topicRepository) Show(req request.TopicShowReq) (topic *response.TopicShowRes, err error) {
	query := r.db.Model(&entity.Topic{})
	query.Where("id = ?", req.TopicId).First(&topic)
	if topic.Id == 0 {
		err = errors.New("未查询到指定主题")
	}
	topic.StatusText = enum.UserMemberIfVerifyMap[topic.Status]
	return
}

func (r *topicRepository) List(req request.TopicListReq) (list response.TopicListRes) {
	var totalNum int64
	query := r.db.Model(&entity.Topic{})
	if req.Title != "" {
		query.Where("title LIKE ?", "%"+req.Title+"%")
	}
	query.Where("status = ?", consts.TopicStatusNormal)
	if req.Order == "hot" {
		query.Order("quote_num desc")
	} else {
		query.Order("created_at desc")
	}
	query.Session(&gorm.Session{}).Count(&totalNum)
	query.Offset(helper.Offset(req.PageReq)).Limit(req.PageReq.PageSize).Scan(&list.List)
	for index, item := range list.List {
		list.List[index].StatusText = enum.TopicStatusMap[item.Status]
	}
	list.Page = response.PageRes{
		PageSize: req.PageReq.PageSize,
		CurrPage: req.PageReq.CurrPage,
		Total:    int(totalNum),
	}
	return
}
