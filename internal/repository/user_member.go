package repository

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/jinzhu/copier"
	"go-start/config"
	"go-start/internal/consts"
	"go-start/internal/model/entity"
	"go-start/internal/model/enum"
	"go-start/internal/pkg/helper"
	"go-start/internal/pkg/jwt"
	"go-start/internal/pkg/mysql"
	"go-start/internal/pkg/redis"
	"go-start/internal/request"
	"go-start/internal/response"
	"go-start/internal/service"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"strconv"
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
	var list []entity.UserMember
	query := r.db.Model(&entity.UserMember{})
	if req.Passport != "" {
		query.Where("passport LIKE ?", "%"+req.Passport+"%")
	}
	query.Offset(helper.Offset(req.PageReq)).Limit(req.PageSize)
	query.Scan(&list)
	return &list
}

func (r *userMemberManageRepository) ChangeStatus(req request.UserMemberChangeStatusReq) error {
	return r.db.Model(&entity.UserMember{}).Where("id = ?", req.UserMemberId).Update("status", req.ToStatus).Error
}

func (r *userMemberRepository) UpdateAvatar(id int64, avatar string) error {
	query := r.db.Model(&entity.UserMember{})
	return query.Where("id = ?", id).Update("avatar", avatar).Error
}

func (r *userMemberRepository) Suggest(req request.UserMemberSuggestReq) (res response.UserMemberSuggestListRes, err error) {
	query := r.db.Model(&entity.UserMember{})
	err = query.Where("nickname LIKE ?", "%"+req.Nickname+"%").Scan(&res.List).Error
	return
}
func (r *userMemberRepository) Show(id int64) (res *response.UserMemberShowRes, err error) {
	query := r.db.Model(&entity.UserMember{})
	query.Where("id = ?", id).First(&res)
	if res.Id == 0 {
		err = errors.New("未查询到指定玩家信息")
	}
	res.StatusText = enum.UserMemberIfVerifyMap[res.Status]
	res.IfVerifyText = enum.UserMemberIfVerifyMap[res.IfVerify]
	return
}

func (r *userMemberRepository) VerifyEmail(req request.UserMemberVerifyEmailReq) error {
	updateMap := map[string]interface{}{"if_verify": consts.UserMemberIfVerifyTrue, "verify_code": "null"}
	return r.db.Model(&entity.UserMember{}).Select([]string{"if_verify", "verify_code"}).Where("id = ? AND verify_code = ?", req.UserMemberId, req.Code).Updates(updateMap).Error
}

func (r *userMemberRepository) GetNotVerify() []entity.UserMember {
	var result []entity.UserMember
	err := r.db.Model(&entity.UserMember{}).Where("if_verify = ? AND verify_code IS NOT NULL", consts.UserMemberIfVerifyFalse).Scan(&result).Error
	if err != nil {
		return []entity.UserMember{}
	}
	return result
}

func (r *userMemberRepository) SignUp(req request.UserMemberSignUpReq) error {
	var (
		um   entity.UserMember
		user response.UserMemberShowRes
	)
	query := r.db.Model(&entity.UserMember{})
	if req.Passport != "" {
		query.Where("passport = ?", req.Passport).First(&user)
	}
	if user.Id != 0 {
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
	um.CreatedAt = helper.FTime{Time: time.Now()}
	um.VerifyCode, _ = helper.String(6, helper.Alphanumeric)
	return r.db.Create(&um).Error
}

func (r *userMemberRepository) SignIn(req request.UserMemberSignInReq) (token string, err error) {
	var (
		user response.UserMemberShowRes
		cfg  = config.Cfg.Jwt
	)
	query := r.db.Model(&entity.UserMember{})
	if req.Passport != "" {
		query.Where("passport = ?", req.Passport).First(&user)
	}
	if user.Id == 0 {
		return "", errors.New("账户不存在")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return "", errors.New("账户与密码不匹配")
	}
	if user.Status == consts.UserMemberStatusForbidden {
		return "", errors.New("该账户已禁用")
	}
	token, err = jwt.Generate(cfg.JwtSecret, cfg.TokenExpire, user, cfg.TokenIssuer)
	if err != nil {
		return "", errors.New("令牌生成失败")
	}
	loginKey := config.Cfg.Redis.LoginPrefix + strconv.Itoa(int(user.Id))
	bs, _ := json.Marshal(user)
	_, err = redis.Client.Set(context.Background(), loginKey, string(bs), 3600*time.Second).Result()
	if err != nil {
		return "", errors.New("用户信息写入缓存失败")
	}
	err = r.db.Model(&entity.UserMember{}).Where("id = ?", user.Id).Update("last_login_ip", req.Ip).Error
	return
}
