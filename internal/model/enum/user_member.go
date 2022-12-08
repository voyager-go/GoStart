package enum

import "go-start/internal/consts"

var (
	UserMemberStatusMap = map[string]string{
		consts.UserMemberStatusNormal:    "正常",
		consts.UserMemberStatusForbidden: "禁用",
	}

	UserMemberIfVerifyMap = map[string]string{
		consts.UserMemberIfVerifyTrue:  "已验证",
		consts.UserMemberIfVerifyFalse: "未验证",
	}
)
