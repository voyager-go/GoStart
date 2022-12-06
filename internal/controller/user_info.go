package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-start/internal/consts/e"
	"go-start/internal/request"
)

type cUser struct {
}

var User = cUser{}

// Create
// @BasePath /api
// @Tags 用户管理
// @Summary 创建用户
// @Schemes
// @Description 新增用户信息
// @Accept json
// @Produce json
// @Param req body request.UserInfoCreateReq true "用户信息"
// @Success 200 {string} c.R.Success
// @Router /user [post]
func (c cUser) Create(ctx *gin.Context) {
	var (
		req request.UserInfoCreateReq
	)
	err := Validator(ctx, &req)
	if err != nil {
		R.Fail(ctx, e.RequestParamsError, err.Error())
		return
	}
	err = DataProvider.UserInfoService.Create(req)
	if err != nil {
		R.Fail(ctx, e.Failed, err.Error())
		return
	}
	R.Success(ctx)
}

// Show
// @BasePath /api
// @Tags 用户管理
// @Summary 查询用户
// @Schemes
// @Description 展示指定用户信息
// @Accept json
// @Produce json
// @Success 200 {object} response.UserInfoShowRes
// @Router /user/show/{id} [get]
func (c cUser) Show(ctx *gin.Context) {
	var (
		req request.UserInfoShowReq
	)
	err := Validator(ctx, &req)
	if err != nil {
		R.Fail(ctx, e.RequestParamsError, err.Error())
		return
	}
	fmt.Println(req)
	res, err := DataProvider.UserInfoService.Show(req)
	R.SuccessWithData(ctx, res)
}
