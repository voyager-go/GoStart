package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	v1 "go-start/api/v1"
	"go-start/internal/consts/e"
	"go-start/internal/pkg/response"
	"go-start/internal/pkg/validator_trans"
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
		UserName: "Hello go-start",
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
