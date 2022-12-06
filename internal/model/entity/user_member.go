package entity

import "time"

type UserMember struct {
	Id          int64     `gorm:"primaryKey"` // 主键编号
	Passport    string    `gorm:"unique"`     // 账户
	Password    string    // 密码
	Nickname    string    // 昵称
	Email       string    `gorm:"unique"` // 邮箱
	LastLoginIp string    // 最后一次登录的IP地址
	Status      int8      // 当前状态 1-启用 2-禁用
	IfVerify    int8      // 是否已经验证 1-已验证 2-未验证
	CreatedAt   time.Time // 创建时间
	UpdatedAt   time.Time // 更新时间
}

func (UserMember) TableName() string {
	return "gs_user_member"
}
