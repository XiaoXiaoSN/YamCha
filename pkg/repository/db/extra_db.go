package db

import (
	"context"
	"yamcha/pkg/model"
)

// GetExtraList get what kind extras be provided in the store
func (repo *dbRepository) GetExtraList(ctx context.Context, branchStoreID int) ([]model.Extra, error) {
	branchStore := model.BranchStore{}
	extraList := []model.Extra{}

	err := repo.db.Model(&model.BranchStore{}).
		Where("id = ?", branchStoreID).
		Find(&branchStore).Error
	if err != nil {
		return nil, err
	}

	err = repo.db.Model(&model.Extra{}).
		Where("store_id = ?", branchStore.StoreGroupID).
		Find(&extraList).Error
	if err != nil {
		return nil, err
	}

	return extraList, nil
}
