package entity

type PostTopic struct {
	BasicEntity
	Uuid    string `gorm:"size:80"` // 帖子唯一标识
	TopicId int64  // 话题编号
}

func (PostTopic) TableName() string {
	return "gs_post_topic"
}
