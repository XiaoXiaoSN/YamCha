package repository

import (
	"context"
	"yamcha/pkg/api/user"

	"github.com/jinzhu/gorm"
)

// UserRepository implement a user Repository
type UserRepository struct {
	db *gorm.DB
}

// NewUserRepository make a user Repository
func NewUserRepository(db *gorm.DB) user.Repository {
	return &UserRepository{
		db: db,
	}
}

// CreateUser ...
func (repo *UserRepository) CreateUser(ctx context.Context, u user.User) error {
	err := repo.db.Model(&user.User{}).Create(&u).Error
	if err != nil {
		return err
	}

	return nil
}

// UserList ...
func (repo *UserRepository) UserList(ctx context.Context) ([]user.User, error) {
	userList := []user.User{}

	err := repo.db.Model(&user.User{}).Find(&userList).Error
	if err != nil {
		return []user.User{}, err
	}

	return userList, nil
}
