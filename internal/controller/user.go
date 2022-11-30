package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go-start/internal/consts/e"
	"go-start/internal/pkg/response"
	"go-start/internal/pkg/validator_trans"
	v1 "go-start/internal/request"
)

type cUser struct {
}

var (
	User = cUser{}
	r    = new(response.R)
)

// @BasePath /api

// Show
// @Summary 展示用户信息
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /user/show/:uuid [get]
func (cUser) Show(ctx *gin.Context) {
	r.SuccessWithData(ctx, struct {
		UserName string
	}{
		UserName: "Hello go-start",
	})
}

// @BasePath /api

// Create
// @Summary 创建用户信息
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /user [post]
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
