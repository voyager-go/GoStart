package response

import (
	"GoStart/internal/consts/e"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var formatResponse map[string]interface{}

// Response 请求返回接口
type Response interface {
	// Success 请求成功
	Success(ctx *gin.Context)
	// Fail 失败请求
	Fail(ctx *gin.Context, code e.Code)
	// SuccessWithData 请求成功并返回数据
	SuccessWithData(ctx *gin.Context, data struct{})
	// FailWithData 请求失败并返回数据
	FailWithData(ctx *gin.Context, data struct{})
	// UnAuthorized 请求未认证
	UnAuthorized(ctx *gin.Context)
	// RequestNotFound 请求不存在
	RequestNotFound(ctx *gin.Context)
	// DefaultRes 默认响应
	DefaultRes(ctx *gin.Context, code e.Code, message string, data interface{})
}

// R 实现Response接口
type R struct{}

func (r *R) DefaultRes(ctx *gin.Context, code e.Code, message string, data interface{}) {
	formatResponse = map[string]interface{}{
		"code":    code,
		"message": message,
		"data":    data,
	}
	ctx.JSON(http.StatusOK, formatResponse)
}

func (r *R) Success(ctx *gin.Context) {
	r.DefaultRes(ctx, e.OK, e.CodeMsg(e.OK), nil)
}

func (r *R) Fail(ctx *gin.Context, code e.Code) {
	r.DefaultRes(ctx, code, e.CodeMsg(e.Failed), nil)
}

func (r *R) SuccessWithData(ctx *gin.Context, data interface{}) {
	r.DefaultRes(ctx, e.OK, e.CodeMsg(e.OK), data)
}

func (r *R) FailWithData(ctx *gin.Context, data interface{}) {
	r.DefaultRes(ctx, e.Failed, e.CodeMsg(e.Failed), data)
}

func (r *R) UnAuthorized(ctx *gin.Context) {
	r.DefaultRes(ctx, e.UnAuthorized, e.CodeMsg(e.UnAuthorized), nil)
}

func (r *R) RequestNotFound(ctx *gin.Context) {
	path := ctx.Request.URL.Path
	method := ctx.Request.Method
	ctx.JSON(http.StatusNotFound, fmt.Sprintf("%s %s not found", method, path))
}
