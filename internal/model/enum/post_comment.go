package enum

import "go-start/internal/consts"

var (
	PostCommentStatusMap = map[int]string{
		consts.PostCommentStatusNormal:    "正常",
		consts.PostCommentStatusForbidden: "禁用",
		consts.PostCommentStatusDeleted:   "删除",
	}
)
