package helper

import (
	"bytes"
	"fmt"
)

// Map2Str 将Map转换为字符串
func Map2Str(m map[string]string) string {
	b := new(bytes.Buffer)
	for k, v := range m {
		fmt.Fprintf(b, "[%s]: %s", k, v)
	}
	return b.String()
}
