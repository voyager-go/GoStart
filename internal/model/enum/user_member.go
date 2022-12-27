package enum

import "go-start/internal/consts"

var (
	UserMemberStatusMap = map[int]string{
		consts.UserMemberStatusNormal:    "正常",
		consts.UserMemberStatusForbidden: "禁用",
	}

	UserMemberIfVerifyMap = map[int]string{
		consts.UserMemberIfVerifyTrue:  "已验证",
		consts.UserMemberIfVerifyFalse: "未验证",
	}
)
