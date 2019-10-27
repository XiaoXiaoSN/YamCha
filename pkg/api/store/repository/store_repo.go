package repository

import (
	"context"
	"yamcha/pkg/api/store"

	"github.com/jinzhu/gorm"
)

// UserRepository implment a user Repository
type StoreRepository struct {
	db *gorm.DB
}

// NewUserRepository make a user Repositoryr
func NewStoreRepository(db *gorm.DB) store.Repository {
	return &StoreRepository{
		db: db,
	}
}

// CreateUser ...
func (repo *StoreRepository) CreateStore(ctx context.Context, u store.Store) (store.Store, error) {
	createItem := u

	err := repo.db.Model(&store.Store{}).Create(&createItem).Error
	if err != nil {
		return u, err
	}

	return createItem, nil
}

// UserList ...
func (repo *StoreRepository) StoreList(ctx context.Context) ([]store.Store, error) {
	storeList := []store.Store{}

	err := repo.db.Model(&store.Store{}).Find(&storeList).Error
	if err != nil {
		return []store.Store{}, err
	}

	return storeList, nil
}
