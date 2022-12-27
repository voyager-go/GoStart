package service

import (
	"go-start/internal/request"
	"go-start/internal/response"
)

type PostManageService interface {
}

type PostService interface {
	Show(req request.PostShowReq) (Post *response.PostShowRes, err error)
	List(req request.PostListReq) (list response.PostListRes)
	Star(req request.PostStarReq) (err error)
	Collect(req request.PostCollectReq) (err error)
	Publish(req request.PostPublishReq) (err error)
	CommentList(req request.PostCommentListReq) (commentList response.PostCommentListRes)
	CommentCreate(req request.PostCommentCreateReq) (err error)
}
