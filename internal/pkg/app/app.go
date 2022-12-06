package app

import (
	"context"
	"encoding/json"
	"errors"
	"go-start/config"
	"go-start/internal/model/entity"
	"go-start/internal/pkg/jwt"
	"go-start/internal/pkg/redis"
	"strconv"
)

type LoginUser struct {
	entity.UserMember
}

type TokenPayload struct {
	UserId int64 `json:"id"`
}

var (
	LoginInvalidErr = "非法授权"
	TokenExpireErr  = "会话过期，请重新授权"
)

// ParseUserByToken 从令牌中解析出授权的用户编号
func ParseUserByToken(token string) (TokenPayload, error) {
	var user TokenPayload
	jwtPayload, err := jwt.Parse(token, config.Cfg.Jwt.JwtSecret)
	if err != nil {
		return user, err
	}
	by, err := json.Marshal(jwtPayload.User)
	if err != nil {
		return user, err
	}
	err = json.Unmarshal(by, &user)
	if err != nil {
		return user, err
	}
	if user.UserId == 0 {
		return user, errors.New(LoginInvalidErr)
	}
	loginKey := config.Cfg.Redis.LoginPrefix + strconv.FormatInt(user.UserId, 10)
	_, err = redis.Client.Get(context.Background(), loginKey).Result()
	if err != nil {
		return TokenPayload{}, errors.New(TokenExpireErr)
	}
	return user, nil
}
