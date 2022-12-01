package request

type PageReq struct {
	PageSize string `json:"page_size"` // 每页条目
	CurrPage string `json:"curr_page"` // 当前页
}
