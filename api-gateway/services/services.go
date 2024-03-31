package services

import (
	"fmt"

	"third-exam/api-gateway/config"
	p "third-exam/api-gateway/genproto/post"
	u "third-exam/api-gateway/genproto/user"
	c "third-exam/api-gateway/genproto/comment"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
)

type IServiceManager interface {
	UserService() u.UserServiceClient
	PostService() p.PostServiceClient
	CommentService() c.CommentServiceClient
}

type serviceManager struct {
	userService    u.UserServiceClient
	postService p.PostServiceClient
	commentService c.CommentServiceClient
}

func (s *serviceManager) UserService() u.UserServiceClient {
	return s.userService
}

func (s *serviceManager) PostService() p.PostServiceClient {
	return s.postService
}

func (s *serviceManager) CommentService() c.CommentServiceClient {
	return s.commentService
}

func NewServiceManager(conf *config.Config) (IServiceManager, error) {
	resolver.SetDefaultScheme("dns")

	connUser, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.UserServiceHost, conf.UserServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	connPost, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.PostServiceHost, conf.PostServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	connComment, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.CommentServiceHost, conf.CommentServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	serviceManager := &serviceManager{
		userService:    u.NewUserServiceClient(connUser),
		postService: p.NewPostServiceClient(connPost),
		commentService: c.NewCommentServiceClient(connComment),
	}

	return serviceManager, nil
}
