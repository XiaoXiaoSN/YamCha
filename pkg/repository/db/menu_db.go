package db

import (
	"context"
	"yamcha/pkg/model"
)

// GetMenuList ...
func (repo *dbRepository) GetMenuList(ctx context.Context, branchStoreID int) ([]model.Menu, error) {
	branchStoreObject := model.BranchStore{}
	menuObject := []model.Menu{}

	errorMsg := repo.db.Model(&model.BranchStore{}).Where("id = ?", branchStoreID).Find(&branchStoreObject).Error
	if errorMsg != nil {
		return []model.Menu{}, errorMsg
	}

	// search everything with store id
	err := repo.db.Model(&model.Menu{}).Where("store_id = ?", branchStoreObject.StoreGroupID).Find(&menuObject).Error
	if err != nil {
		return []model.Menu{}, err
	}

	return menuObject, nil

}
