package request

type PageReq struct {
	PageSize int `json:"size" form:"size" uri:"size" binding:"required,numeric"` // 每页条目
	CurrPage int `json:"page" form:"page" uri:"page" binding:"required,numeric"` // 当前页
}
