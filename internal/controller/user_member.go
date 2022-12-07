package controller

import (
	"github.com/gin-gonic/gin"
	"go-start/internal/consts/e"
	"go-start/internal/pkg/log"
	"go-start/internal/request"
)

type cUserMember struct {
}

var UserMember = cUserMember{}

// SignUp
// @BasePath /api
// @Tags 玩家信息
// @Summary 注册
// @Schemes
// @Description 填写基本信息
// @Accept json
// @Produce json
// @Param req body request.UserMemberSignUpReq true "玩家信息"
// @Success 200 {string} c.R.Success
// @Router /user-member/sign-up [post]
func (cUserMember) SignUp(ctx *gin.Context) {
	var (
		req request.UserMemberSignUpReq
	)
	err := Validator(ctx, &req)
	if err != nil {
		R.Fail(ctx, e.RequestParamsError, err.Error())
		return
	}
	err = DataProvider.UserMemberService.SignUp(req)
	if err != nil {
		R.Fail(ctx, e.Failed, err.Error())
		return
	}
	R.Success(ctx)
}

// SignIn
// @BasePath /api
// @Tags 玩家信息
// @Summary 登录
// @Schemes
// @Description 填写登录信息
// @Accept json
// @Produce json
// @Param req body request.UserMemberSignInReq true "登录信息"
// @Success 200 {string} c.R.Success
// @Router /user-member/sign-in [post]
func (cUserMember) SignIn(ctx *gin.Context) {
	var (
		req request.UserMemberSignInReq
	)
	err := Validator(ctx, &req)
	if err != nil {
		R.Fail(ctx, e.RequestParamsError, err.Error())
		return
	}
	req.Ip = ctx.ClientIP()
	token, err := DataProvider.UserMemberService.SignIn(req)
	if err != nil {
		R.Fail(ctx, e.Failed, err.Error())
		return
	}
	R.SuccessWithData(ctx, token)
}

// Show
// @BasePath /api
// @Tags 玩家信息
// @Summary 查询
// @Schemes
// @Description 填写基本信息
// @Accept json
// @Produce json
// @Param user_member_id path request.UserMemberShowReq true "玩家主键编号"
// @Success 200 {string} c.R.Success
// @Router /user-member/{user_member_id} [get]
func (cUserMember) Show(ctx *gin.Context) {
	var (
		req request.UserMemberShowReq
	)
	err := Validator(ctx, &req)
	if err != nil {
		R.Fail(ctx, e.RequestParamsError, err.Error())
		return
	}
	res, err := DataProvider.UserMemberService.Show(req)
	log.Logger.Infof("resdata %v", res)
	if err != nil {
		R.Fail(ctx, e.Failed, err.Error())
		return
	}
	R.SuccessWithData(ctx, res)
}
