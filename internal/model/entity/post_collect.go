package entity

type PostCollect struct {
	BasicEntity
	Uuid         string `gorm:"size:80"` // 帖子唯一标识
	UserMemberId int64  // 收藏人
}

func (PostCollect) TableName() string {
	return "gs_post_collect"
}
