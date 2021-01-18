package db

import (
	"context"
	"log"
	"yamcha/pkg/model"
)

// CreateStore ...
func (repo *dbRepository) CreateStore(ctx context.Context, targetStore model.Store) (model.Store, error) {
	err := repo.db.Model(&model.Store{}).Create(&targetStore).Error
	if err != nil {
		return targetStore, err
	}

	return targetStore, nil
}

// GetStore ...
func (repo *dbRepository) GetStore(ctx context.Context, storeID int) (model.Store, error) {
	targetStore := model.Store{}

	err := repo.db.Model(&model.Store{}).Preload("BranchStores").Find(&targetStore).Error
	if err != nil {
		// TODO: handle mysql 1602 duplicate error (should return http status 404)
		return model.Store{}, err
	}

	return targetStore, nil
}

// StoreList ...
func (repo *dbRepository) StoreList(ctx context.Context) ([]model.Store, error) {
	storeList := []model.Store{}

	err := repo.db.Model(&model.Store{}).Find(&storeList).Error
	if err != nil {
		// log.Printf("Error: %s", err)

		return []model.Store{}, err
	}

	return storeList, nil
}

// BranchStoreList ...
func (repo *dbRepository) BranchStoreList(ctx context.Context, storeID int) ([]model.BranchStore, error) {
	log.Println("in Branch:", storeID)
	branchStoreList := []model.BranchStore{}

	err := repo.db.Model(&model.BranchStore{}).Where("store_group_id = ?", storeID).Find(&branchStoreList).Error
	log.Println("in Branch:", branchStoreList)
	if err != nil {
		return []model.BranchStore{}, err
	}

	return branchStoreList, nil
}

// CreateBranchStore ...
func (repo *dbRepository) CreateBranchStore(ctx context.Context, branchStore model.BranchStore) (model.BranchStore, error) {
	err := repo.db.Model(&model.Store{}).Create(&branchStore).Error
	if err != nil {
		return branchStore, err
	}

	return branchStore, nil
}
