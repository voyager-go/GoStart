package request

type UserCreateReq struct {
	UserName string `json:"user_name" binding:"required"`
}
