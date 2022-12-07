package entity

import "go-start/internal/pkg/helper"

type BasicEntity struct {
	Id        int64        `gorm:"primaryKey"` // 主键编号
	CreatedAt helper.FTime // 创建时间
	UpdatedAt helper.FTime // 更新时间
}
