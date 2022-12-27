package request

// UserMemberSignUpReq 玩家注册输入参数
type UserMemberSignUpReq struct {
	Passport string `json:"passport" binding:"required,min=2,max=30"`
	Password string `json:"password" binding:"required,min=6,max=20"`
	Email    string `json:"email" binding:"required,email"`
	Nickname string `json:"nickname" binding:"omitempty,min=2,max=30"`
}

// UserMemberSignInReq 玩家登录输入参数
type UserMemberSignInReq struct {
	Passport string `json:"passport" binding:"required,min=2,max=30"`
	Password string `json:"password" binding:"required,min=6,max=20"`
	Ip       string `json:"-"`
}

// UserMemberListReq 玩家列表输入参数
type UserMemberListReq struct {
	Passport string `json:"passport" binding:"omitempty"`
	PageReq
}

// UserMemberChangeStatusReq 玩家请求修改状态输入参数
type UserMemberChangeStatusReq struct {
	UserMemberId int64  `json:"user_member_id" binding:"required"`
	ToStatus     string `json:"to_status" binding:"required"`
}

// UserMemberSuggestReq 根据昵称查找玩家列表输入参数
type UserMemberSuggestReq struct {
	Nickname string `json:"nickname" form:"nickname" binding:"required"`
}

// UserMemberVerifyEmailReq 邮箱验证
type UserMemberVerifyEmailReq struct {
	UserMemberId int64  `form:"uid" uri:"uid" binding:"required"`
	Code         string `form:"code" uri:"code" binding:"required"`
}
