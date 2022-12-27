package helper

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"go-start/internal/request"
	"os"
	"reflect"
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
func FormatTimestamp(currTime time.Time, layout string) string {
	if layout == "" {
		layout = "2006-01-02 15:04:05"
	}
	return currTime.Format(layout)
}

// Offset 分页偏移量计算
func Offset(req request.PageReq) int {
	if req.CurrPage < 1 {
		req.CurrPage = 1
	}
	return (req.CurrPage - 1) * req.PageSize
}

// InArray 判断字符是否在数组中
func InArray(obj interface{}, target interface{}) bool {
	targetValue := reflect.ValueOf(target)
	switch reflect.TypeOf(target).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < targetValue.Len(); i++ {
			if targetValue.Index(i).Interface() == obj {
				return true
			}
		}
	case reflect.Map:
		if targetValue.MapIndex(reflect.ValueOf(obj)).IsValid() {
			return true
		}
	}

	return false
}

// IsFileExist 判断文件或者文件夹是否存在
func IsFileExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// Md5 MD5摘要
func Md5(str string) string {
	data := []byte(str)
	m := md5.New()
	hashStr := m.Sum(data)
	return fmt.Sprintf("%x", hashStr)
}
