package request

type UserInfoShowReq struct {
	Id       string `json:"id" binding:"omitempty"`
	Passport string `json:"passport" binding:"omitempty,min=2,max=30"`
}

type UserInfoCreateReq struct {
	Passport string `json:"passport" binding:"required,min=2,max=30"`
	Password string `json:"password" binding:"required,min=6,max=20"`
	Nickname string `json:"nickname" binding:"omitempty,min=2,max=30"`
}

type UserInfoListReq struct {
	Passport string `json:"passport" binding:"omitempty,string"`
	PageReq
}
