package service

import (
	"context"
	c "third-exam/comment-service/genproto/comment"
	l "third-exam/comment-service/pkg/logger"
	grpcClient "third-exam/comment-service/service/grpc_client"
	storage "third-exam/comment-service/storage"

	"github.com/jmoiron/sqlx"
)

// CommentService ...
type CommentService struct {
	storage storage.IStorage
	logger  l.Logger
	client  grpcClient.IServiceManager
}

// NewCommentService ...
func NewCommentService(db *sqlx.DB, log l.Logger, client grpcClient.IServiceManager) *CommentService {
	return &CommentService{
		storage: storage.NewStoragePg(db),
		logger:  log,
		client:  client,
	}
}

// Create implements CommentService.
func (s *CommentService) CreateComment(ctx context.Context, req *c.Comment) (*c.Comment, error) {
	pro, err := s.storage.Comment().CreateComment(req)
	if err != nil {
		return nil, err
	}
	return pro, nil
}

// Get implements CommentService
func (s *CommentService) GetComment(ctx context.Context, req *c.GetCommentRequest) (*c.Comment, error) {

	pro, err := s.storage.Comment().GetComment(req)
	if err != nil {
		s.logger.Error(err.Error())
	}

	return pro, nil
}

// GetAllComment implements comment.CommentServiceServer.
func (s *CommentService) GetAllComment(ctx context.Context, req *c.GetAllCommentRequest) (*c.GetAllCommentResponse, error) {
	res, err := s.storage.Comment().GetAllComments(req)

	if err != nil {
		return nil, err
	}

	return res, nil
}

// Update implements product.ProductServiceServer.
func (s *CommentService) UpdateComment(ctx context.Context, req *c.Comment) (*c.Comment, error) {
	pro, err := s.storage.Comment().UpdateComment(req)
	if err != nil {
		return nil, err
	}
	return pro, nil
}

// Delete implements PostService.
func (s *CommentService) DeleteComment(ctx context.Context, req *c.GetDeleteCommentRequest) (*c.Comment, error) {
	pro, err := s.storage.Comment().DeleteComment(req)
	if err != nil {
		return nil, err
	}
	return pro, nil
}
