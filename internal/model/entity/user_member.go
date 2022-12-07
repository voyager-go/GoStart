package entity

type UserMember struct {
	BasicEntity
	Passport    string `gorm:"unique"` // 账户
	Password    string // 密码
	Nickname    string // 昵称
	Email       string `gorm:"unique"` // 邮箱
	LastLoginIp string // 最后一次登录的IP地址
	Status      int8   `gorm:"default:1"` // 当前状态 1-启用 2-禁用
	IfVerify    int8   `gorm:"default:2"` // 是否已经验证 1-已验证 2-未验证
}

func (UserMember) TableName() string {
	return "gs_user_member"
}
