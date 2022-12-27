package entity

type PostStar struct {
	BasicEntity
	Uuid         string `gorm:"size:80"` // 帖子唯一标识
	UserMemberId int64  // 点赞人
}

func (PostStar) TableName() string {
	return "gs_post_star"
}
