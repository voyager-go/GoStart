package entity

type UserInfo struct {
	BasicEntity
	Passport string `gorm:"unique"` // 账户
	Password string // 密码
	Nickname string // 昵称
}

func (UserInfo) TableName() string {
	return "gs_user_info"
}
