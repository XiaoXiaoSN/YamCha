package repository

import (
	"context"
	"yamcha/pkg/api/user"

	"github.com/jinzhu/gorm"
)

// UserRepository implment a user Repository
type UserRepository struct {
	db *gorm.DB
}

// NewUserRepository make a user Repositoryr
func NewUserRepository(db *gorm.DB) user.Repository {
	return &UserRepository{
		db: db,
	}
}

// CreateUser ...
func (svc *UserRepository) CreateUser(ctx context.Context, u user.User) error {
	return nil
}
