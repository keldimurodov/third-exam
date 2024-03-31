package repo

import (
	c "third-exam/comment-service/genproto/comment"
)

// PostStorageI ...
type CommentStorageI interface {
	CreateComment(c *c.Comment) (*c.Comment, error)
	GetComment(req *c.GetCommentRequest) (*c.Comment, error)
	GetAllComments(req *c.GetAllCommentRequest) (*c.GetAllCommentResponse, error)
	UpdateComment(*c.UpdateRequest) (*c.Comment, error)
	DeleteComment(*c.GetDeleteCommentRequest) (*c.Comment, error)
}
