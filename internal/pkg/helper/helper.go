package helper

import (
	"bytes"
	"fmt"
	"go-start/internal/request"
	"strconv"
	"time"
)

// Map2Str 将Map转换为字符串
func Map2Str(m map[string]string) string {
	b := new(bytes.Buffer)
	for k, v := range m {
		fmt.Fprintf(b, "[%s]: %s", k, v)
	}
	return b.String()
}

// FormatTimestamp 格式化时间输出
func FormatTimestamp(currTime time.Time) string {
	layout := "2006-01-02 15:04:05"
	return currTime.Format(layout)
}

// Offset 分页偏移量计算
func Offset(req request.PageReq) int {
	currPage, _ := strconv.Atoi(req.CurrPage)
	perPage, _ := strconv.Atoi(req.PageSize)
	if currPage < 1 {
		currPage = 1
	}
	return (currPage - 1) * perPage
}
