package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go-start/internal/pkg/log"
	"time"
)

func LoggerToFile() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//开始时间
		startTime := time.Now()
		//处理请求
		ctx.Next()
		//结束时间
		endTime := time.Now()
		//执行时间
		latencyTime := endTime.Sub(startTime)

		//请求方式
		reqMethod := ctx.Request.Method

		//请求路由
		reqUri := ctx.Request.RequestURI

		//状态码
		statusCode := ctx.Writer.Status()

		//请求IP
		ClientIp := ctx.ClientIP()

		//日志格式
		log.Logger.WithFields(logrus.Fields{
			"status_code":  statusCode,
			"latency_time": latencyTime,
			"client_ip":    ClientIp,
			"req_method":   reqMethod,
			"req_uri":      reqUri,
		}).Info()
	}
}
