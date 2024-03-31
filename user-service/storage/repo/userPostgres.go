package repo

import (
	u "third-exam/user-service/genproto/user"
)

// UserStorageI ...
type UserStoragePostgresI interface {
	CreateUser(*u.User) (*u.User, error)
	GetUser(user *u.GetUserRequest) (*u.User, error)
	GetAllUsers(req *u.GetAllRequest) (*u.GetAllResponse, error)
	DeleteUser(user *u.GetUserRequest) (*u.User, error)
	UpdateUser(user *u.User) (*u.User, error)
	CheckUniqueness(req *u.CheckUniquenessRequest) (*u.CheckUniquenessResponse, error)
	Login(req *u.LoginRequest) (*u.User, error)
}
