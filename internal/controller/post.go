package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-start/internal/consts/e"
	"go-start/internal/request"
)

type cPost struct {
}

var Post = cPost{}

// Show
// @BasePath /api
// @Tags 主题
// @Summary 查询
// @Schemes
// @Description 查询主题信息
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization"
// @Param uuid query string true  "推文uuid" request.PostShowReq
// @Success 200 {string} c.R.Success
// @Router /post/show [get]
func (cPost) Show(ctx *gin.Context) {
	var (
		req request.PostShowReq
	)
	err := Validator(ctx, &req)
	if err != nil {
		R.Fail(ctx, e.RequestParamsError, err.Error())
		return
	}
	fmt.Println(req)
	res, err := DataProvider.PostService.Show(req)
	if err != nil {
		R.Fail(ctx, e.Failed, err.Error())
		return
	}
	R.SuccessWithData(ctx, res)
}

// List
// @BasePath /api
// @Tags 推文
// @Summary 查询
// @Schemes
// @Description 查询主题信息
// @Accept json
// @Produce json
// @Param list query string true  "列表输入参数" request.PostListReq
// @Success 200 {string} c.R.Success
// @Router /post/list [get]
func (cPost) List(ctx *gin.Context) {
	var (
		req request.PostListReq
	)
	err := Validator(ctx, &req)
	if err != nil {
		R.Fail(ctx, e.RequestParamsError, err.Error())
		return
	}
	list := DataProvider.PostService.List(req)
	if err != nil {
		R.Fail(ctx, e.Failed, err.Error())
		return
	}
	R.SuccessWithData(ctx, list)
}

// CommentList
// @BasePath /api
// @Tags 推文
// @Summary 查询评论
// @Schemes
// @Description 查询推文的评论列表
// @Accept json
// @Produce json
// @Param list query string true  "评论列表输入参数" request.PostCommentListReq
// @Success 200 {string} c.R.Success
// @Router /post/comment-list [get]
func (cPost) CommentList(ctx *gin.Context) {
	var (
		req request.PostCommentListReq
	)
	err := Validator(ctx, &req)
	if err != nil {
		R.Fail(ctx, e.RequestParamsError, err.Error())
		return
	}
	list := DataProvider.PostService.CommentList(req)
	if err != nil {
		R.Fail(ctx, e.Failed, err.Error())
		return
	}
	R.SuccessWithData(ctx, list)
}

// CommentCreate
// @BasePath /api
// @Tags 推文
// @Summary 评论
// @Schemes
// @Description 评论推文或者回复评论
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization"
// @Param comment body request.PostCommentCreateReq true "评论信息"
// @Success 200 {string} c.R.Success
// @Router /post/publish [post]
func (cPost) CommentCreate(ctx *gin.Context) {
	var (
		req request.PostCommentCreateReq
	)
	err := Validator(ctx, &req)
	if err != nil {
		R.Fail(ctx, e.RequestParamsError, err.Error())
		return
	}
	req.UserMemberId, err = GetUidByToken(ctx)
	if err != nil {
		R.Fail(ctx, e.Failed, err.Error())
		return
	}
	req.Ip = ctx.ClientIP()
	err = DataProvider.PostService.CommentCreate(req)
	if err != nil {
		R.Fail(ctx, e.Failed, err.Error())
		return
	}
	R.Success(ctx)
}

// Publish
// @BasePath /api
// @Tags 推文
// @Summary 发布
// @Schemes
// @Description 发布一条推文
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization"
// @Param publish_info body request.PostPublishReq true "推文信息"
// @Success 200 {string} c.R.Success
// @Router /post/publish [post]
func (cPost) Publish(ctx *gin.Context) {
	var (
		req request.PostPublishReq
	)
	err := Validator(ctx, &req)
	if err != nil {
		R.Fail(ctx, e.RequestParamsError, err.Error())
		return
	}
	req.UserMemberId, err = GetUidByToken(ctx)
	if err != nil {
		R.Fail(ctx, e.Failed, err.Error())
		return
	}
	req.PublishedIp = ctx.ClientIP()
	err = DataProvider.PostService.Publish(req)
	if err != nil {
		R.Fail(ctx, e.Failed, err.Error())
		return
	}
	R.Success(ctx)
}

// Star
// @BasePath /api
// @Tags 推文
// @Summary 点赞
// @Schemes
// @Description 给推文点赞
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization"
// @Param uuid body request.PostStarReq true "推文编号"
// @Success 200 {string} c.R.Success
// @Router /post/star [post]
func (cPost) Star(ctx *gin.Context) {
	var (
		req request.PostStarReq
	)
	err := Validator(ctx, &req)
	if err != nil {
		R.Fail(ctx, e.RequestParamsError, err.Error())
		return
	}
	req.UserMemberId, err = GetUidByToken(ctx)
	if err != nil {
		R.Fail(ctx, e.Failed, err.Error())
		return
	}
	err = DataProvider.PostService.Star(req)
	if err != nil {
		R.Fail(ctx, e.Failed, err.Error())
		return
	}
	R.Success(ctx)
}

// Collect
// @BasePath /api
// @Tags 推文
// @Summary 收藏
// @Schemes
// @Description 收藏推文
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization"
// @Param uuid body request.PostCollectReq true "推文编号"
// @Success 200 {string} c.R.Success
// @Router /post/collect [post]
func (cPost) Collect(ctx *gin.Context) {
	var (
		req request.PostCollectReq
	)
	err := Validator(ctx, &req)
	if err != nil {
		R.Fail(ctx, e.RequestParamsError, err.Error())
		return
	}
	req.UserMemberId, err = GetUidByToken(ctx)
	if err != nil {
		R.Fail(ctx, e.Failed, err.Error())
		return
	}
	err = DataProvider.PostService.Collect(req)
	if err != nil {
		R.Fail(ctx, e.Failed, err.Error())
		return
	}
	R.Success(ctx)
}
