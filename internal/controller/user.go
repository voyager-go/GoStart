package controller

import (
	v1 "GoStart/api/v1"
	"GoStart/internal/consts/e"
	"GoStart/internal/pkg/response"
	"GoStart/internal/pkg/validator_trans"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type cUser struct {
}

var (
	User = cUser{}
	r    = new(response.R)
)

func (cUser) Show(ctx *gin.Context) {
	r.SuccessWithData(ctx, struct {
		UserName string
	}{
		UserName: "Hello GoStart",
	})
}

func (cUser) Create(ctx *gin.Context) {
	var (
		req v1.UserCreateReq
	)
	if err := ctx.ShouldBindJSON(&req); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			r.Fail(ctx, e.RequestParamsError, err.Error())
			return
		}
		r.Fail(ctx, e.RequestParamsError, validator_trans.Translate(errs)[0])
		return
	}
	r.SuccessWithData(ctx, struct {
		UserName string
	}{
		UserName: req.UserName,
	})
}
