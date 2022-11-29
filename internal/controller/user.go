package controller

import (
	"GoStart/internal/pkg/response"
	"github.com/gin-gonic/gin"
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
	r.SuccessWithData(ctx, struct {
		UserName string
	}{
		UserName: ctx.PostForm("user_name"),
	})
}
