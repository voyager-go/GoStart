package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-start/config"
	"go-start/internal/consts/e"
	"go-start/internal/pkg/helper"
	"go-start/internal/request"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

type cUserMember struct {
}

var UserMember = cUserMember{}

func (cUserMember) VerifyEmail(ctx *gin.Context) {
	var req request.UserMemberVerifyEmailReq
	err := Validator(ctx, &req)
	if err != nil {
		R.Fail(ctx, e.RequestParamsError, err.Error())
		return
	}
	err = DataProvider.UserMemberService.VerifyEmail(req)
	if err != nil {
		R.Fail(ctx, e.Failed, err.Error())
		return
	}
	R.Success(ctx)
}

// UploadAvatar
// @BasePath /api
// @Tags 玩家信息
// @Summary 上传头像
// @Schemes
// @Description 上传用户的头像
// @Accept json
// @Produce json
// @Param avatar body string true "头像信息"
// @Success 200 {string} c.R.Success
// @Router /user-member/upload-avatar [post]
func (cUserMember) UploadAvatar(ctx *gin.Context) {
	f, err := ctx.FormFile("avatar")
	if err != nil {
		R.Fail(ctx, e.Failed, "未获取到文件信息")
		return
	}
	fExt := strings.ToLower(path.Ext(f.Filename))
	fmt.Println(f.Filename, fExt)
	if !helper.InArray(fExt, []string{".png", ".jpg", ".jpeg"}) {
		R.Fail(ctx, e.Failed, "仅支持以png|jpeg|jpg后缀结尾的图片")
		return
	}
	userMemberId, _ := GetUidByToken(ctx)
	fName := "avatar/" + helper.FormatTimestamp(time.Now(), "2006-01-02") + "/"
	fPath := config.Cfg.Server.UploadPath + fName
	isExist, err := helper.IsFileExist(fPath)
	if !isExist {
		err := os.MkdirAll(fPath, os.ModePerm)
		if err != nil {
			R.Fail(ctx, e.Failed, err.Error())
			return
		}
		// 创建文件夹的同时创建一个html文件方式文件夹被列出
		input, _ := os.ReadFile(config.Cfg.Server.UploadPath + "index.html")
		err = os.WriteFile(fPath+"index.html", input, os.ModePerm)
		if err != nil {
			R.Fail(ctx, e.Failed, err.Error())
			return
		}
	}
	filename := fPath + helper.Md5(strconv.FormatInt(userMemberId, 10)) + fExt
	if err := ctx.SaveUploadedFile(f, filename); err != nil {
		R.Fail(ctx, e.Failed, err.Error())
		return
	}
	R.SuccessWithData(ctx, fName+helper.Md5(strconv.FormatInt(userMemberId, 10))+fExt)
}

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
// @Success 200 {string} c.R.Success
// @Router /user-member/show [get]
func (cUserMember) Show(ctx *gin.Context) {
	userMemberId, _ := GetUidByToken(ctx)
	res, err := DataProvider.UserMemberService.Show(userMemberId)
	if err != nil {
		R.Fail(ctx, e.Failed, err.Error())
		return
	}
	R.SuccessWithData(ctx, res)
}

// Suggest
// @BasePath /api
// @Tags 玩家信息
// @Summary 根据用户名联想的玩家
// @Schemes
// @Description 填写基本信息
// @Accept json
// @Produce json
// @Success 200 {string} c.R.Success
// @Router /user-member/suggest [get]
func (cUserMember) Suggest(ctx *gin.Context) {
	var (
		req request.UserMemberSuggestReq
	)
	err := Validator(ctx, &req)
	if err != nil {
		R.Fail(ctx, e.RequestParamsError, err.Error())
		return
	}
	res, err := DataProvider.UserMemberService.Suggest(req)
	fmt.Println(res)
	if err != nil {
		R.Fail(ctx, e.Failed, err.Error())
		return
	}
	R.SuccessWithData(ctx, res)
}

//func (cUserMember) List(ctx *gin.Context) {
//	var (
//		req request.UserMemberListReq
//	)
//	err := Validator(ctx, &req)
//	if err != nil {
//		R.Fail(ctx, e.RequestParamsError, err.Error())
//		return
//	}
//	res := DataProvider.UserMemberManageService.List(req)
//	if err != nil {
//		R.Fail(ctx, e.Failed, err.Error())
//		return
//	}
//	R.SuccessWithData(ctx, res)
//}
