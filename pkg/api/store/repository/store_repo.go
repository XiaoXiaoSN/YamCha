package repository

import (
	"context"
	"yamcha/pkg/api/store"

	"github.com/jinzhu/gorm"
)

// StoreRepository implment a user Repository
type StoreRepository struct {
	db *gorm.DB
}

// NewStoreRepository make a user Repositoryr
func NewStoreRepository(db *gorm.DB) store.Repository {
	return &StoreRepository{
		db: db,
	}
}

// CreateStore ...
func (repo *StoreRepository) CreateStore(ctx context.Context, u store.Store) (store.Store, error) {
	createItem := u

	err := repo.db.Model(&store.Store{}).Create(&createItem).Error
	if err != nil {
		return u, err
	}

	return createItem, nil
}

// StoreList ...
func (repo *StoreRepository) StoreList(ctx context.Context) ([]store.Store, error) {
	storeList := []store.Store{}

	err := repo.db.Model(&store.Store{}).Find(&storeList).Error
	if err != nil {
		return []store.Store{}, err
	}

	return storeList, nil
}

// BranchStoreList ...
func (repo *StoreRepository) BranchStoreList(ctx context.Context, id string) ([]store.BranchStore, error) {
	branchStoreList := []store.BranchStore{}

	err := repo.db.Model(&store.BranchStore{}).Where("store_group_id = ?", id).Find(&branchStoreList).Error
	if err != nil {
		return []store.BranchStore{}, err
	}

	return branchStoreList, nil
}
