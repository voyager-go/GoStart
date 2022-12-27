package request

type PostCommentListReq struct {
	Uuid    string `json:"uuid" form:"uuid" uri:"uuid" binding:"required"` // 帖子唯一标识
	PageReq `binding:"required"`
}

type PostCommentCreateReq struct {
	Uuid         string `json:"uuid" binding:"required"`    // 帖子唯一标识
	UserMemberId int64  `json:"-"`                          // 评论人
	Pid          *int64 `json:"pid" binding:"required"`     // 上层用户ID，如果是0，表示评论的是帖子
	Content      string `json:"content" binding:"required"` // 评论内容
	Ip           string `json:"-"`                          // 评论的IP地址
}
