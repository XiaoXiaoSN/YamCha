package db

import (
	"context"
	"yamcha/pkg/model"
)

// GetExtraList ...
func (repo *dbRepository) GetExtraList(ctx context.Context, branchStoreID int) ([]model.Extra, error) {
	extraArray := []model.Extra{}
	branchStoreObject := model.BranchStore{}

	errorMsg := repo.db.Model(&model.BranchStore{}).
		Where("id = ?", branchStoreID).
		Find(&branchStoreObject).Error
	if errorMsg != nil {
		return []model.Extra{}, errorMsg
	}

	// search everything with store id
	err := repo.db.Model(&model.Extra{}).
		Where("store_id = ?", branchStoreObject.StoreGroupID).
		Find(&extraArray).Error
	if err != nil {
		return []model.Extra{}, err
	}

	return extraArray, nil
}
