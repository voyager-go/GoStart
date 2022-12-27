package response

import "go-start/internal/pkg/helper"

type PostCommentShowRes struct {
	UserMember   UserMember   `json:"user_member" gorm:"-"` // 评论用户信息
	ToMember     UserMember   `json:"to_member" gorm:"-"`   // 回复用户信息
	Id           int64        `json:"id"`                   // 评论编号
	Uuid         string       `json:"uuid"`                 // 帖子唯一标识
	UserMemberId int          `json:"user_member_id"`       // 评论人
	Pid          int64        `json:"pid"`                  // 上层用户ID，如果是0，表示评论的是帖子
	Content      string       `json:"content"`              // 评论内容
	Status       int          `json:"status"`               // 当前状态 1-正常 2-封禁 3-删除
	StatusText   string       `json:"status_text" gorm:"-"`
	Ip           string       `json:"ip"` // 评论的IP地址
	City         string       `json:"city"`
	CreatedAt    helper.FTime `json:"created_at"`
	UpdatedAt    helper.FTime `json:"updated_at"`
}

type PostCommentListRes struct {
	List  []PostCommentShowRes `json:"list"`  // 列表信息
	Pager PageRes              `json:"pager"` // 分页信息
}
