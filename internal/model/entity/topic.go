package entity

type Topic struct {
	BasicEntity
	Title    string `gorm:"unique,size:200"` // 话题名称
	Icon     string `gorm:"size:300"`        // 话题小图标
	QuoteNum int64  // 引用次数
	Status   int8   `gorm:"default:1"` // 话题状态 1-正常 2-封禁 3-删除
}

func (Topic) TableName() string {
	return "gs_topic"
}
