package service

import (
	"context"
	pb "third-exam/user-service/genproto/post"
	l "third-exam/user-service/pkg/logger"
	grpcClient "third-exam/user-service/service/grpc_client"
	storage "third-exam/user-service/storage"

	"github.com/jmoiron/sqlx"
)

// PostService ...
type PostService struct {
	storage storage.IStorage
	logger  l.Logger
	client  grpcClient.IServiceManager
}

// NewPostService ...
func NewPostService(db *sqlx.DB, log l.Logger, client grpcClient.IServiceManager) *PostService {
	return &PostService{
		storage: storage.NewStoragePg(db),
		logger:  log,
		client:  client,
	}
}

// Create implements PostService.
func (s *PostService) CreatePost(ctx context.Context, req *pb.Post) (*pb.Post, error) {
	pro, err := s.storage.Post().CreatePost(req)
	if err != nil {
		return nil, err
	}
	return pro, nil
}

// Get implements PostService
func (s *PostService) GetPost(ctx context.Context, req *pb.GetPostRequest) (*pb.Post, error) {

	pro, err := s.storage.Post().GetPost(req)
	if err != nil {
		s.logger.Error(err.Error())
	}

	return pro, nil
}

// GetAll implements PostService
func (s *PostService) GetAllPosts(ctx context.Context, req *pb.GetAllRequest) (*pb.GetAllResponse, error) {
	med, err := s.storage.Post().GetAllPosts(req)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}

	return med, nil
}

// Update implements product.ProductServiceServer.
func (s *PostService) UpdatePost(ctx context.Context, req *pb.Post) (*pb.Post, error) {
	pro, err := s.storage.Post().UpdatePost(req)
	if err != nil {
		return nil, err
	}
	return pro, nil
}

// Delete implements PostService.
func (s *PostService) DeletePost(ctx context.Context, req *pb.GetDeletePostRequest) (*pb.Post, error) {
	pro, err := s.storage.Post().DeletePost(req)
	if err != nil {
		return nil, err
	}
	return pro, nil
}
