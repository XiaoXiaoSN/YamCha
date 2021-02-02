package db

import (
	"context"
	"yamcha/pkg/model"
)

// CreateStore ...
func (repo *dbRepository) CreateStore(ctx context.Context, targetStore *model.Store) error {
	err := repo.db.Model(&model.Store{}).Create(targetStore).Error
	if err != nil {
		return err
	}
	return nil
}

// GetStore ...
func (repo *dbRepository) GetStore(ctx context.Context, storeID int) (model.Store, error) {
	targetStore := model.Store{}

	db := repo.db.Model(&model.Store{})
	db = db.Preload("BranchStores")

	err := db.Where("id = ?", storeID).
		First(&targetStore).Error
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
		return []model.Store{}, err
	}

	return storeList, nil
}

// BranchStoreList ...
func (repo *dbRepository) BranchStoreList(ctx context.Context, storeID int) ([]model.BranchStore, error) {
	branchStoreList := []model.BranchStore{}

	err := repo.db.Model(&model.BranchStore{}).
		Where("store_group_id = ?", storeID).
		Find(&branchStoreList).Error
	if err != nil {
		return []model.BranchStore{}, err
	}

	return branchStoreList, nil
}

// CreateBranchStore ...
func (repo *dbRepository) CreateBranchStore(ctx context.Context, branchStore *model.BranchStore) error {
	err := repo.db.Model(&model.BranchStore{}).
		Create(branchStore).Error
	if err != nil {
		return err
	}
	return nil
}
