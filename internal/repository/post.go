package repository

import (
	"errors"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"github.com/yinheli/qqwry"
	"go-start/config"
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
	_ service.PostService       = (*postRepository)(nil)
	_ service.PostManageService = (*postManageRepository)(nil)
)

type postManageRepository struct {
	db *gorm.DB
}

type postRepository struct {
	db *gorm.DB
}

func newPostRepository() service.PostService {
	return &postRepository{db: mysql.Conn}
}

func newPostManageRepository() service.PostManageService {
	return &postManageRepository{db: mysql.Conn}
}

func (r *postRepository) Publish(req request.PostPublishReq) (err error) {
	ipPackage := qqwry.NewQQwry(config.Cfg.Server.QqwryPath)
	ipPackage.Find(req.PublishedIp)
	var post = entity.Post{
		Uuid:          uuid.NewV4().String(),
		UserMemberId:  req.UserMemberId,
		Content:       req.Content,
		PublishedCity: ipPackage.Country,
		PublishedIp:   req.PublishedIp,
		IfTop:         1,
	}
	err = r.db.Model(&entity.Post{}).Create(&post).Error
	return
}

func (r *postRepository) CommentCreate(req request.PostCommentCreateReq) (err error) {
	ipPackage := qqwry.NewQQwry(config.Cfg.Server.QqwryPath)
	ipPackage.Find(req.Ip)
	var comment = entity.PostComment{
		Uuid:         req.Uuid,
		UserMemberId: req.UserMemberId,
		Pid:          *req.Pid,
		Content:      req.Content,
		Status:       consts.PostCommentStatusNormal,
		Ip:           req.Ip,
		City:         ipPackage.Country,
	}
	return r.db.Model(&entity.PostComment{}).Create(&comment).Error
}

func (r *postRepository) Show(req request.PostShowReq) (post *response.PostShowRes, err error) {
	query := r.db.Model(&entity.Post{})
	query.Where("uuid = ?", req.Uuid).First(&post)
	if post.Uuid == "" {
		err = errors.New("未查询到指定主题")
	}
	post.StatusText = enum.UserMemberIfVerifyMap[post.Status]
	post.IfTopText = enum.UserMemberIfVerifyMap[post.IfTop]
	post.IfVisibleText = enum.PostIfVisibleMap[post.IfVisible]

	var member response.UserMemberShowRes
	r.db.Model(&entity.UserMember{}).Where("id = ?", post.UserMemberId).First(&member)
	post.UserMember = response.UserMember{
		Id:       member.Id,
		Avatar:   member.Avatar,
		Passport: member.Passport,
		Nickname: member.Nickname,
	}
	return
}

func (r *postRepository) CommentList(req request.PostCommentListReq) (commentList response.PostCommentListRes) {

	var totalNum int64
	query := r.db.Model(&entity.PostComment{})
	query.Where("status = ? AND uuid = ?", consts.PostCommentStatusNormal, req.Uuid)
	query.Session(&gorm.Session{}).Count(&totalNum)
	query.Order("created_at desc")
	query.Offset(helper.Offset(req.PageReq)).Limit(req.PageReq.PageSize).Scan(&commentList.List)
	for index, item := range commentList.List {
		commentList.List[index].StatusText = enum.PostCommentStatusMap[item.Status]
		var (
			member   response.UserMemberShowRes
			toMember response.UserMemberShowRes
		)
		r.db.Model(&entity.UserMember{}).Where("id = ?", item.UserMemberId).First(&member)
		commentList.List[index].UserMember = response.UserMember{
			Id:       member.Id,
			Avatar:   member.Avatar,
			Passport: member.Passport,
			Nickname: member.Nickname,
		}
		if item.Pid != 0 {
			r.db.Model(&entity.UserMember{}).Where("id = ?", item.Pid).First(&toMember)
			commentList.List[index].ToMember = response.UserMember{
				Id:       member.Id,
				Avatar:   member.Avatar,
				Passport: member.Passport,
				Nickname: member.Nickname,
			}
		}
	}
	commentList.Pager = response.PageRes{
		PageSize: req.PageReq.PageSize,
		CurrPage: req.PageReq.CurrPage,
		Total:    int(totalNum),
	}
	return
}

func (r *postRepository) List(req request.PostListReq) (list response.PostListRes) {

	var totalNum int64
	query := r.db.Model(&entity.Post{})
	if req.Keywords != "" {
		query.Where("content LIKE ?", "%"+req.Keywords+"%")
	}
	if req.Passport != "" {
		var userMember entity.UserMember
		r.db.Model(&entity.UserMember{}).Where("passport = ?", req.Passport).Find(&userMember)
		query.Where("user_member_id = ?", userMember.Id)
	}
	query.Where("status = ?", consts.PostStatusNormal)
	query.Session(&gorm.Session{}).Count(&totalNum)
	query.Order("created_at desc,collect_num desc,comment_num desc,star_num desc")
	query.Offset(helper.Offset(req.PageReq)).Limit(req.PageReq.PageSize).Scan(&list.List)
	for index, item := range list.List {
		list.List[index].StatusText = enum.PostStatusMap[item.Status]
		list.List[index].IfTopText = enum.PostStatusMap[item.IfTop]
		list.List[index].IfVisibleText = enum.PostIfVisibleMap[item.IfVisible]
		var member response.UserMemberShowRes
		r.db.Model(&entity.UserMember{}).Where("id = ?", item.UserMemberId).First(&member)
		list.List[index].UserMember = response.UserMember{
			Id:       member.Id,
			Avatar:   member.Avatar,
			Passport: member.Passport,
			Nickname: member.Nickname,
		}
	}
	list.Pager = response.PageRes{
		PageSize: req.PageReq.PageSize,
		CurrPage: req.PageReq.CurrPage,
		Total:    int(totalNum),
	}
	return
}

func (r *postRepository) Star(req request.PostStarReq) (err error) {

	var (
		data = entity.PostStar{
			Uuid:         req.Uuid,
			UserMemberId: req.UserMemberId,
		}
		condition = "uuid = ? AND user_member_id = ?"
		operate   = "+"
		count     int64
	)
	r.db.Model(&entity.PostStar{}).Where(condition, data.Uuid, data.UserMemberId).Count(&count)
	if req.Operate == "sub" {
		operate = "-"
		if count <= 0 {
			return errors.New("操作异常")
		}
	} else {
		if count > 0 {
			return errors.New("无法重复点赞")
		}
	}

	if err = r.db.Transaction(func(tx *gorm.DB) error {
		if operate == "+" {
			if err := tx.Model(&entity.PostStar{}).Create(&data).Error; err != nil {
				return err
			}
		} else {
			if err := tx.Where(condition, data.Uuid, data.UserMemberId).Delete(&entity.PostStar{}).Error; err != nil {
				return err
			}
		}
		expr := fmt.Sprintf("star_num %s ?", operate)
		if err := tx.Model(&entity.Post{}).Where("uuid = ?", data.Uuid).UpdateColumn("star_num", gorm.Expr(expr, 1)).Error; err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}
	return
}

func (r *postRepository) Collect(req request.PostCollectReq) (err error) {
	var (
		data = entity.PostCollect{
			Uuid:         req.Uuid,
			UserMemberId: req.UserMemberId,
		}
		condition = "uuid = ? AND user_member_id = ?"
		operate   = "+"
		count     int64
	)
	r.db.Model(&entity.PostCollect{}).Where(condition, data.Uuid, data.UserMemberId).Count(&count)
	if req.Operate == "sub" {
		operate = "-"
		if count <= 0 {
			return errors.New("操作异常")
		}
	} else {
		if count > 0 {
			return errors.New("无法重复收藏")
		}
	}

	if err = r.db.Transaction(func(tx *gorm.DB) error {
		if operate == "+" {
			if err := tx.Model(&entity.PostCollect{}).Create(&data).Error; err != nil {
				return err
			}
		} else {
			if err := tx.Where(condition, data.Uuid, data.UserMemberId).Delete(&entity.PostCollect{}).Error; err != nil {
				return err
			}
		}
		expr := fmt.Sprintf("collect_num %s ?", operate)
		if err := tx.Model(&entity.Post{}).Where("uuid = ?", data.Uuid).UpdateColumn("collect_num", gorm.Expr(expr, 1)).Error; err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}
	return
}
