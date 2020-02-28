package repository

import (
	"context"
	"log"
	"yamcha/pkg/api/store"

	"github.com/jinzhu/gorm"
)

// StoreRepository implement a user Repository
type StoreRepository struct {
	db *gorm.DB
}

// NewStoreRepository make a user Repository
func NewStoreRepository(db *gorm.DB) store.Repository {
	return &StoreRepository{
		db: db,
	}
}

// CreateStore ...
func (repo *StoreRepository) CreateStore(ctx context.Context, targetStore store.Store) (store.Store, error) {
	err := repo.db.Model(&store.Store{}).Create(&targetStore).Error
	if err != nil {
		return targetStore, err
	}

	return targetStore, nil
}

// GetStore ...
func (repo *StoreRepository) GetStore(ctx context.Context, storeID int) (store.Store, error) {
	targetStore := store.Store{}

	err := repo.db.Model(&store.Store{}).Preload("BranchStores").Find(&targetStore).Error
	if err != nil {
		// TODO: handle mysql 1602 duplicate error (should return http status 404)
		return store.Store{}, err
	}

	return targetStore, nil
}

// StoreList ...
func (repo *StoreRepository) StoreList(ctx context.Context) ([]store.Store, error) {
	storeList := []store.Store{}

	err := repo.db.Model(&store.Store{}).Find(&storeList).Error
	if err != nil {
		// log.Printf("Error: %s", err)

		return []store.Store{}, err
	}

	return storeList, nil
}

// BranchStoreList ...
func (repo *StoreRepository) BranchStoreList(ctx context.Context, storeID int) ([]store.BranchStore, error) {
	log.Println("in Branch:", storeID)
	branchStoreList := []store.BranchStore{}

	err := repo.db.Model(&store.BranchStore{}).Where("store_group_id = ?", storeID).Find(&branchStoreList).Error
	log.Println("in Branch:", branchStoreList)
	if err != nil {
		return []store.BranchStore{}, err
	}

	return branchStoreList, nil
}

// CreateBranchStore ...
func (repo *StoreRepository) CreateBranchStore(ctx context.Context, branchStore store.BranchStore) (store.BranchStore, error) {
	err := repo.db.Model(&store.Store{}).Create(&branchStore).Error
	if err != nil {
		return branchStore, err
	}

	return branchStore, nil
}
