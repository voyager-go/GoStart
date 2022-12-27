package response

type PageRes struct {
	PageSize int `json:"size"`  // 每页条目
	CurrPage int `json:"page"`  // 当前页
	Total    int `json:"total"` // 总条目
}
