package service

import (
	"context"
	"yamcha/pkg/api/user"
)

// UserService implement a user service
type UserService struct {
	UserRepo user.Repository
}

// NewUserService make a user servicer
func NewUserService(userRepo user.Repository) user.Service {
	return &UserService{
		UserRepo: userRepo,
	}
}

// CreateUser ...
func (svc *UserService) CreateUser(ctx context.Context, u user.User) error {
	return svc.UserRepo.CreateUser(ctx, u)
}

// UserList ...
func (svc *UserService) UserList(ctx context.Context) ([]user.User, error) {
	return svc.UserRepo.UserList(ctx)
}
