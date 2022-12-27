package entity

type PostComment struct {
	BasicEntity
	Uuid         string `gorm:"size:80"` // 帖子唯一标识
	UserMemberId int64  // 评论人
	Pid          int64  // 上层用户ID，如果是0，表示评论的是帖子
	Content      string `gorm:"size:300"`  // 评论内容
	Status       int8   `gorm:"default:1"` // 当前状态 1-正常 2-封禁 3-删除
	Ip           string // 评论的IP地址
	City         string `gorm:"size:50"` // 评论的城市
}

func (PostComment) TableName() string {
	return "gs_post_comment"
}
