package repository

import (
	"errors"
	"github.com/jinzhu/copier"
	"go-start/internal/model/entity"
	"go-start/internal/pkg/mysql"
	"go-start/internal/request"
	"go-start/internal/response"
	"go-start/internal/service"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

var (
	_ service.UserMemberService       = (*userMemberRepository)(nil)
	_ service.UserMemberManageService = (*userMemberManageRepository)(nil)
)

type userMemberRepository struct {
	db *gorm.DB
}

type userMemberManageRepository struct {
	db *gorm.DB
}

func newUserMemberRepository() service.UserMemberService {
	return &userMemberRepository{db: mysql.Conn}
}

func newUserMemberManageRepository() service.UserMemberManageService {
	return &userMemberManageRepository{db: mysql.Conn}
}

func (r *userMemberManageRepository) List(req request.UserMemberListReq) *[]entity.UserMember {
	return nil
}

func (r *userMemberManageRepository) ChangeStatus(req request.UserMemberChangeStatusReq) error {
	return nil
}

func (r *userMemberRepository) Show(req request.UserMemberShowReq) (res *response.UserMemberShowRes, err error) {
	return
}
func (r *userMemberRepository) SignUp(req request.UserMemberSignUpReq) error {
	var (
		um   entity.UserMember
		user response.UserMemberShowRes
	)
	if req.Passport == "" {
		return errors.New("查询条件缺失")
	}
	query := r.db.Model(&entity.UserMember{})
	if req.Passport != "" {
		query.Where("passport = ?", req.Passport).First(&user)
	}
	if user.Id != "" {
		return errors.New("passport 已存在")
	}
	if err := copier.Copy(&um, req); err != nil {
		return err
	}
	by, err := bcrypt.GenerateFromPassword([]byte(um.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	um.Password = string(by)
	um.CreatedAt = time.Now()
	return r.db.Create(&user).Error
}

func (r *userMemberRepository) SignIn(req request.UserMemberSignInReq) string {
	return ""
}
