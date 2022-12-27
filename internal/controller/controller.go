package controller

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go-start/config"
	"go-start/internal/pkg/app"
	"go-start/internal/pkg/res"
	"go-start/internal/pkg/validator_trans"
	"go-start/internal/repository"
	"strings"
)

var (
	// R 请求返回
	R *res.R
	// DataProvider 服务提供者
	DataProvider *repository.DataService
	// Validator JSON参数请求验证器
	Validator     func(ctx *gin.Context, req interface{}) (err error)
	GetUidByToken func(ctx *gin.Context) (int64, error)
)

func InitMethods() {
	R = new(res.R)
	DataProvider = repository.NewDataProvider()
	Validator = func(ctx *gin.Context, req interface{}) (err error) {
		err = ctx.ShouldBind(req)
		fmt.Println(req)
		if err != nil {
			errs, ok := err.(validator.ValidationErrors)
			if ok {
				err = errors.New(validator_trans.Translate(errs)[0])
			}
		}
		return
	}

	GetUidByToken = func(ctx *gin.Context) (int64, error) {
		token := ctx.GetHeader(config.Cfg.Jwt.TokenKey)
		if strings.HasPrefix(token, "Bearer") {
			token = strings.TrimPrefix(token, "Bearer ")
		}
		t, err := app.ParseUserByToken(token)
		if err != nil {
			return 0, err
		}
		return t.UserId, nil
	}
}
