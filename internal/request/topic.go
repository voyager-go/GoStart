package request

// TopicShowReq 主题信息输入参数
type TopicShowReq struct {
	TopicId int64 `json:"topic_id" uri:"topic_id" binding:"required,numeric"`
}

// TopicListReq 主题列表输入参数
type TopicListReq struct {
	Title   string `json:"title" form:"title" uri:"title" binding:"omitempty"`
	Order   string `json:"order" form:"order" uri:"order" binding:"oneof=hot new"`
	PageReq `binding:"required"`
}

// TopicChangeStatusReq 主题请求修改状态输入参数
type TopicChangeStatusReq struct {
	TopicId  int64  `json:"topic_id" binding:"required,numeric"`
	ToStatus string `json:"to_status" binding:"required"`
}
