package enum

import "go-start/internal/consts"

var (
	TopicStatusMap = map[int]string{
		consts.TopicStatusNormal:    "正常",
		consts.TopicStatusForbidden: "禁用",
		consts.TopicStatusDeleted:   "删除",
	}
)
