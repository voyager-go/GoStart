package response

import "go-start/internal/pkg/helper"

type PostShowRes struct {
	UserMember    `json:"user_member"  gorm:"-"`
	Uuid          string       `json:"uuid"`
	UserMemberId  int64        `json:"user_member_id"`
	Content       string       `json:"content"`
	StarNum       string       `json:"star_num"`
	CommentNum    string       `json:"comment_num"`
	CollectNum    string       `json:"collect_num"`
	PublishedCity string       `json:"published_city"`
	IfTop         int          `json:"if_top"`
	IfTopText     string       `json:"if_top_text"  gorm:"-"`
	IfVisible     int          `json:"if_visible"`
	IfVisibleText string       `json:"if_visible_text"  gorm:"-"`
	Status        int          `json:"status"`
	StatusText    string       `json:"status_text"  gorm:"-"` // 话题状态 1-正常 2-封禁 3-删除
	CreatedAt     helper.FTime `json:"created_at"`
	UpdatedAt     helper.FTime `json:"updated_at"`
}

type PostListRes struct {
	List  []PostShowRes `json:"list"`  // 列表信息
	Pager PageRes       `json:"pager"` // 分页信息
}
