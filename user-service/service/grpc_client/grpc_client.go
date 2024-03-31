package grpcClient

import (
	"fmt"
	config "third-exam/user-service/config"
	c "third-exam/user-service/genproto/comment"
	p "third-exam/user-service/genproto/post"

	"google.golang.org/grpc"
)

type IServiceManager interface {
	PostService() p.PostServiceClient
	CommentService() c.CommentServiceClient
}
type serviceManager struct {
	cfg     config.Config
	Post    p.PostServiceClient
	Comment c.CommentServiceClient
}

func (s *serviceManager) PostService() p.PostServiceClient {
	return s.Post
}

func (s *serviceManager) CommentService() c.CommentServiceClient {
	return s.Comment
}

func New(cfg config.Config) (IServiceManager, error) {
	PostConnection, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.PostServiceHost, cfg.PostServicePort),
		grpc.WithInsecure())

	if err != nil {
		return nil, err
	}

	CommentConnection, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.CommentsServiceHost, cfg.CommentsServicePort),
		grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return &serviceManager{
		cfg:     cfg,
		Post:    p.NewPostServiceClient(PostConnection),
		Comment: c.NewCommentServiceClient(CommentConnection)}, nil

}



// // GrpcClientI ...
// type GrpcClientI interface {
// }

// // GrpcClient ...
// type GrpcClient struct {
// 	cfg         config.Config
// 	connections map[string]interface{}
// }

// // New ...
// func New(cfg config.Config) (*GrpcClient, error) {
// 	return &GrpcClient{
// 		cfg:         cfg,
// 		connections: map[string]interface{}{},
// 	}, nil
// }
