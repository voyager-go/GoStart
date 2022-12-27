package request

type PostCollectReq struct {
	Uuid         string `json:"uuid" binding:"required"`                  // 推文编号
	UserMemberId int64  `json:"-"`                                        // 创建人
	Operate      string `json:"operate" binding:"required,oneof=add sub"` // add 是收藏 sub是取消收藏
}
