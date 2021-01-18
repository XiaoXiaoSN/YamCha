package service

import (
	"context"
	"yamcha/pkg/api/user"
	"yamcha/pkg/model"
	"yamcha/pkg/repository"
)

// UserService implement a user service
type UserService struct {
	repo repository.Repository
}

// NewUserService make a user servicer
func NewUserService(repo repository.Repository) user.Service {
	return &UserService{
		repo: repo,
	}
}

// CreateUser ...
func (svc *UserService) CreateUser(ctx context.Context, u model.User) error {
	return svc.repo.CreateUser(ctx, u)
}

// UserList ...
func (svc *UserService) UserList(ctx context.Context) ([]model.User, error) {
	return svc.repo.UserList(ctx)
}
