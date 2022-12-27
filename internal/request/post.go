package request

// PostShowReq 主题信息输入参数
type PostShowReq struct {
	Uuid string `json:"uuid" form:"uuid" uri:"uuid" binding:"required"`
}

// PostListReq 主题列表输入参数
type PostListReq struct {
	Keywords string `json:"keywords" form:"keywords" uri:"keywords" binding:"omitempty"`
	Passport string `json:"passport" form:"passport" uri:"passport" binding:"omitempty"`
	PageReq  `binding:"required"`
}

// PostChangeStatusReq 主题请求修改状态输入参数
type PostChangeStatusReq struct {
	Uuid     string `json:"uuid" binding:"required"`
	ToStatus string `json:"to_status" binding:"required"`
}

type PostPublishReq struct {
	UserMemberId  int64  `json:"user_member_id"`             // 创建人
	Content       string `json:"content" binding:"required"` // 内容
	PublishedCity string `json:"published_city"`             // 发布城市
	PublishedIp   string `json:"published_ip"`               // 发布Ip
	IfTop         string `json:"if_top"`                     // 是否置顶 1-是 2-否
}
