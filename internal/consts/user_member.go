package consts

const (
	UserMemberStatusNormal    = 1 // 玩家账户正常
	UserMemberStatusForbidden = 2 // 玩家账户封禁
	UserMemberIfVerifyTrue    = 1 // 玩家已验证
	UserMemberIfVerifyFalse   = 2 // 玩家未验证
	UserMemberVerifyLink      = "http://127.0.0.1:8080/api/user-member/verify-email?uid=%d&code=%s"
)
