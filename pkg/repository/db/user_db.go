package db

import (
	"context"
	"yamcha/pkg/model"
)

// CreateUser ...
func (repo *dbRepository) CreateUser(ctx context.Context, u *model.User) error {
	err := repo.db.Model(&model.User{}).Create(u).Error
	if err != nil {
		return err
	}

	return nil
}

// UserList ...
func (repo *dbRepository) UserList(ctx context.Context) ([]model.User, error) {
	userList := []model.User{}

	err := repo.db.Model(&model.User{}).Find(&userList).Error
	if err != nil {
		return []model.User{}, err
	}

	return userList, nil
}
