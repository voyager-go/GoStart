package entity

import (
	"time"
)

type UserInfo struct {
	Id        int64     `gorm:"primaryKey"` // 主键编号
	Passport  string    `gorm:"unique"`     // 账户
	Password  string    // 密码
	Nickname  string    // 昵称
	CreatedAt time.Time // 创建时间
	UpdatedAt time.Time // 更新时间
}

func (UserInfo) TableName() string {
	return "gs_user_info"
}
