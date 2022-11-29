package e

type Code int

const (
	OK     Code = 0
	Failed Code = 10000 + iota
	UnAuthorized
	AuthExpired
	InternalError
	RequestMethodError
	RequestParamsError
)

var (
	codeMap = map[Code]string{
		OK:                 "请求成功",
		Failed:             "请求失败",
		UnAuthorized:       "用户未认证",
		AuthExpired:        "会话过期，请重新登录",
		InternalError:      "服务器内部错误",
		RequestMethodError: "请求方式错误",
		RequestParamsError: "请求参数错误",
	}

	CodeMsg = func(c Code) string {
		if v, ok := codeMap[c]; ok {
			return v
		}
		return "Unknown Response Code"
	}
)
