package grpcClient

import (
	"fmt"
	config "third-exam/post-service/config"
	c "third-exam/post-service/genproto/comment"
	u "third-exam/post-service/genproto/user"

	"google.golang.org/grpc"
)

type IServiceManager interface {
	UserService() u.UserServiceClient
	CommentService() c.CommentServiceClient
}
type serviceManager struct {
	cfg     config.Config
	User    u.UserServiceClient
	Comment c.CommentServiceClient
}

func (s *serviceManager) UserService() u.UserServiceClient {
	return s.User
}

func (s *serviceManager) CommentService() c.CommentServiceClient {
	return s.Comment
}

func New(cfg config.Config) (IServiceManager, error) {
	UserConnection, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.UserServiceHost, cfg.UserServicePort),
		grpc.WithInsecure())

	if err != nil {
		return nil, err
	}

	CommentConnection, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.CommentServiceHost, cfg.CommentServicePort),
		grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return &serviceManager{
		cfg:     cfg,
		User:    u.NewUserServiceClient(UserConnection),
		Comment: c.NewCommentServiceClient(CommentConnection)}, nil

}
