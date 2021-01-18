package user

import (
	"context"
	"yamcha/pkg/model"
)

// Service is a user service
type Service interface {
	UserList(ctx context.Context) ([]model.User, error)
	CreateUser(ctx context.Context, u model.User) error
}
