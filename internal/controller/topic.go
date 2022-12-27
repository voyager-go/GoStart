package controller

import (
	"github.com/gin-gonic/gin"
	"go-start/internal/consts/e"
	"go-start/internal/request"
)

type cTopic struct {
}

var Topic = cTopic{}

// Show
// @BasePath /api
// @Tags 主题
// @Summary 查询
// @Schemes
// @Description 查询主题信息
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization"
// @Param topic_id query int true "主题编号" request.TopicShowReq
// @Success 200 {string} c.R.Success
// @Router /topic/{topic_id} [get]
func (cTopic) Show(ctx *gin.Context) {
	var (
		req request.TopicShowReq
	)
	err := Validator(ctx, &req)
	if err != nil {
		R.Fail(ctx, e.RequestParamsError, err.Error())
		return
	}
	res, err := DataProvider.TopicService.Show(req)
	if err != nil {
		R.Fail(ctx, e.Failed, err.Error())
		return
	}
	R.SuccessWithData(ctx, res)
}

// List
// @BasePath /api
// @Tags 主题
// @Summary 查询
// @Schemes
// @Description 查询主题信息
// @Accept json
// @Produce json
// @Param list body request.TopicListReq true "列表输入参数"
// @Success 200 {string} c.R.Success
// @Router /topic/list [post]
func (cTopic) List(ctx *gin.Context) {
	var (
		req request.TopicListReq
	)
	err := Validator(ctx, &req)
	if err != nil {
		R.Fail(ctx, e.RequestParamsError, err.Error())
		return
	}
	list := DataProvider.TopicService.List(req)
	if err != nil {
		R.Fail(ctx, e.Failed, err.Error())
		return
	}
	R.SuccessWithData(ctx, list)
}
