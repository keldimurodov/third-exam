package service

import (
	"context"
	"database/sql"
	"fmt"
	pbu "third-exam/user-service/genproto/user"
	l "third-exam/user-service/pkg/logger"
	"third-exam/user-service/storage"

	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

// UserService ...
type UserService struct {
	storage      storage.IStorage
	storageRedis storage.IStorageRedis
	logger       l.Logger
	db           *sql.DB
	rdb          *redis.Client
	// rdb redis.Client
}

// NewUserService ...
func NewUserService(db *sqlx.DB, rdb *redis.Client, log l.Logger) *UserService {
	return &UserService{
		storage:      storage.NewStoragePg(db),
		storageRedis: storage.NewStorageRedis(rdb),
		logger:       log,
	}
}

func (s *UserService) CreateUser(ctx context.Context, req *pbu.User) (*pbu.User, error) {
	user, err := s.storage.User().CreateUser(req)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) GetUser(ctx context.Context, req *pbu.GetUserRequest) (*pbu.User, error) {
	user, err := s.storage.User().GetUser(req)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) UpdateUser(ctx context.Context, req *pbu.User) (*pbu.User, error) {
	user, err := s.storage.User().UpdateUser(req)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) DeleteUser(ctx context.Context, req *pbu.GetUserRequest) (*pbu.User, error) {
	user, err := s.storage.User().DeleteUser(req)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) GetAllUsers(ctx context.Context, req *pbu.GetAllRequest) (*pbu.GetAllResponse, error) {
	users, err := s.storage.User().GetAllUsers(req)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *UserService) CheckUniqueness(ctx context.Context, req *pbu.CheckUniquenessRequest) (*pbu.CheckUniquenessResponse, error) {
	user, err := s.storage.User().CheckUniqueness(req)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) Sign(ctx context.Context, user *pbu.UserDetail) (*pbu.ResponseMessage, error) {
	req, err := s.storageRedis.UserRedis().Sign(user)
	if err != nil {
		fmt.Println("big service error")
		return nil, err

	}
	return req, nil
}

func (s *UserService) Verification(ctx context.Context, req *pbu.VerificationUserRequest) (*pbu.User, error) {
	user, err := s.storageRedis.UserRedis().Verification(req)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) Login(ctx context.Context, req *pbu.LoginRequest) (*pbu.User, error) {
	user, err := s.storage.User().Login(req)
	if err != nil {
		return nil, err
	}
	return user, nil
}
