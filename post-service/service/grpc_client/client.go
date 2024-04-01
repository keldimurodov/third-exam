package grpcClient

import (
	"fmt"
	"third-exam/post-service/config"
	c "third-exam/post-service/genproto/comment"
	u "third-exam/post-service/genproto/user"

	"google.golang.org/grpc"
)

type IServiceManager interface {
	UserService() u.UserServiceClient
	CommentService() c.CommentServiceClient
}

type serviceManager struct {
	cfg            config.Config
	userService    u.UserServiceClient
	commentService c.CommentServiceClient
}

func New(cfg config.Config) (IServiceManager, error) {
	// dail to user-service
	connUser, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.UserServiceHost, cfg.UserServicePort),
		grpc.WithInsecure(),
	)
	if err != nil {
		return nil, fmt.Errorf("user service dail host: %s port : %d", cfg.UserServiceHost, cfg.UserServicePort)
	}
	// dail to comment-service
	connComment, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.CommentServiceHost, cfg.CommentServicePort),
		grpc.WithInsecure(),
	)
	if err != nil {
		return nil, fmt.Errorf("comment service dail host: %s port : %d", cfg.CommentServiceHost, cfg.CommentServicePort)
	}
	return &serviceManager{
		cfg:            cfg,
		userService:    u.NewUserServiceClient(connUser),
		commentService: c.NewCommentServiceClient(connComment),
	}, nil
}

func (s *serviceManager) UserService() u.UserServiceClient {
	return s.userService
}

func (s *serviceManager) CommentService() c.CommentServiceClient {
	return s.commentService
}
