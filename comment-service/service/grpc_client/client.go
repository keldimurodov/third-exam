package grpcClient

import (
	"fmt"
	config "third-exam/comment-service/config"
	p "third-exam/comment-service/genproto/post"
	u "third-exam/comment-service/genproto/user"

	"google.golang.org/grpc"
)

type IServiceManager interface {
	UserService() u.UserServiceClient
	PostService() p.PostServiceClient
}
type serviceManager struct {
	cfg  config.Config
	User u.UserServiceClient
	Post p.PostServiceClient
}

func (s *serviceManager) UserService() u.UserServiceClient {
	return s.User
}

func (s *serviceManager) PostService() p.PostServiceClient {
	return s.Post
}

func New(cfg config.Config) (IServiceManager, error) {
	UserConnection, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.UserServiceHost, cfg.UserServicePort),
		grpc.WithInsecure())

	if err != nil {
		return nil, err
	}

	PostConnection, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.PostServiceHost, cfg.PostServicePort),
		grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return &serviceManager{
		cfg:  cfg,
		User: u.NewUserServiceClient(UserConnection),
		Post: p.NewPostServiceClient(PostConnection)}, nil

}
