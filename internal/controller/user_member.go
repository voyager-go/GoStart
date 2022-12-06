package controller

import (
	"github.com/gin-gonic/gin"
	"go-start/internal/consts/e"
	"go-start/internal/request"
)

type cUserMember struct {
}

var UserMember = cUserMember{}

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
