package service

import (
	"context"
	p "third-exam/post-service/genproto/post"
	l "third-exam/post-service/pkg/logger"
	grpcClient "third-exam/post-service/service/grpc_client"
	"third-exam/post-service/storage"

	"github.com/jmoiron/sqlx"
	"go.mongodb.org/mongo-driver/mongo"
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

func (s *PostService) CreatePost(ctx context.Context, req *p.Post) (*p.Post, error) {
	post, err := s.storage.Post().CreatePost(req)
	if err != nil {
		return nil, err
	}

	return post, nil
}

func (s *PostService) UpdatePost(ctx context.Context, req *p.Post) (*p.Post, error) {
	post, err := s.storage.Post().UpdatePost(req)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}

	return post, nil
}

func (s *PostService) DeletePost(ctx context.Context, req *p.GetDeletePostRequest) (*p.Post, error) {
	post, err := s.storage.Post().DeletePost(req)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}

	return post, nil
}

// Get implements PostService
func (s *PostService) GetPost(ctx context.Context, req *p.GetPostRequest) (*p.GetP, error) {

	pro, err := s.storage.Post().GetPost(req)
	if err != nil {
		s.logger.Error(err.Error())
	}

	return pro, nil

}

// GetAll implements PostService
func (s *PostService) GetAllPosts(ctx context.Context, req *p.GetAllRequest) (*p.GetAllResponse, error) {
	med, err := s.storage.Post().GetAllPosts(req)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}

	return med, nil
}

func NewCommentServiceMongo(db *mongo.Collection, log l.Logger, client grpcClient.IServiceManager) *PostService {
	return &PostService{
		storage: storage.NewStorageMongo(db),
		logger:  log,
		client:  client,
	}
}
