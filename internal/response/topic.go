package response

import "go-start/internal/pkg/helper"

type TopicShowRes struct {
	Id         int64        `json:"id"`                    // 编号
	Title      string       `json:"title"`                 // 话题名称
	Icon       string       `json:"icon"`                  // 话题小图标
	QuoteNum   int          `json:"quote_num"`             // 热度
	IfTop      int          `json:"if_top"`                // 是否置顶 1-是 2-否
	IfTopText  string       `json:"if_top_text"`           // 是否置顶 1-是 2-否
	Status     int          `json:"status"`                // 话题状态 1-正常 2-封禁 3-删除
	StatusText string       `json:"status_text"  gorm:"-"` // 话题状态 1-正常 2-封禁 3-删除
	CreatedAt  helper.FTime `json:"created_at"`
	UpdatedAt  helper.FTime `json:"updated_at"`
}

type TopicListRes struct {
	List []TopicShowRes `json:"list"` // 列表信息
	Page PageRes        `json:"page"` // 分页信息
}
