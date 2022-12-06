package request

// UserMemberSignUpReq 玩家注册输入参数
type UserMemberSignUpReq struct {
	Passport string `json:"passport" binding:"required,min=2,max=30"`
	Password string `json:"password" binding:"required,min=6,max=20"`
	Nickname string `json:"nickname" binding:"omitempty,min=2,max=30"`
}

// UserMemberSignInReq 玩家登录输入参数
type UserMemberSignInReq struct {
	Passport string `json:"passport" binding:"required,min=2,max=30"`
	Password string `json:"password" binding:"required,min=6,max=20"`
}

// UserMemberShowReq 玩家信息输入参数
type UserMemberShowReq struct {
	UserMemberId string `json:"user_member_id" binding:"required"`
}

// UserMemberListReq 玩家列表输入参数
type UserMemberListReq struct {
	Passport string `json:"passport" binding:"omitempty,string"`
	PageReq
}

// UserMemberChangeStatusReq 玩家请求修改状态输入参数
type UserMemberChangeStatusReq struct {
	UserMemberId string `json:"user_member_id" binding:"required"`
	ToStatus     string `json:"to_status" binding:"required"`
}
