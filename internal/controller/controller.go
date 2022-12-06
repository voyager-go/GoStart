package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go-start/internal/pkg/res"
	"go-start/internal/pkg/validator_trans"
	"go-start/internal/repository"
)

var (
	// R 请求返回
	R *res.R
	// DataProvider 服务提供者
	DataProvider = repository.NewDataProvider()
	// Validator JSON参数请求验证器
	Validator = func(ctx *gin.Context, req interface{}) (err error) {
		if err = ctx.ShouldBindJSON(&req); err != nil {
			errs, ok := err.(validator.ValidationErrors)
			if ok {
				err = errors.New(validator_trans.Translate(errs)[0])
			}
		}
		return
	}
)
