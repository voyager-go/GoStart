package v1

type UserCreateReq struct {
	UserName string `json:"user_name" binding:"required"`
}
