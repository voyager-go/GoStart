package entity

type Post struct {
	BasicEntity
	Uuid          string `gorm:"unique,size:80"` // 帖子唯一标识
	UserMemberId  int64  // 创建人
	Content       string // 内容
	StarNum       int64  `gorm:"default:0"`  // 点赞数
	CommentNum    int64  `gorm:"default:0"`  // 评论数
	CollectNum    int64  `gorm:"default:0"`  // 收藏数
	PublishedCity string `gorm:"size:50"`    // 发布城市
	PublishedIp   string `gorm:"size:20,ip"` // 发布IP
	IfTop         int8   `gorm:"default:2"`  // 是否置顶 1-是 2-否
	IfVisible     int8   `gorm:"default:1"`  // 是否可见 1-全部可见 2-仅好友可见 3-仅自己可见
	Status        int8   `gorm:"default:1"`  // 当前状态 1-正常 2-封禁 3-删除
}

func (Post) TableName() string {
	return "gs_post"
}
