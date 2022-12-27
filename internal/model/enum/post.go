package enum

import "go-start/internal/consts"

var (
	PostStatusMap = map[int]string{
		consts.PostStatusNormal:    "正常",
		consts.PostStatusForbidden: "禁用",
		consts.PostStatusDeleted:   "删除",
	}
	PostIfTopMap = map[int]string{
		consts.PostIfTopTrue:  "置顶",
		consts.PostIfTopFalse: "不置顶",
	}
	PostIfVisibleMap = map[int]string{
		consts.PostIfVisiblePublic:    "公开可见",
		consts.PostIfVisibleProtected: "仅朋友可见",
		consts.PostIfVisiblePrivate:   "仅自己可见",
	}
)
