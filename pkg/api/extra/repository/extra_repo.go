package repository

import (
	"context"
	"yamcha/pkg/api/extra"
	"yamcha/pkg/api/menu"

	"github.com/jinzhu/gorm"
)

// ExtraRepository implment a extra Repository
type ExtraRepository struct {
	db *gorm.DB
}

// NewExtraRepository make a extra Repositoryr
func NewExtraRepository(db *gorm.DB) extra.Repository {
	return &ExtraRepository{
		db: db,
	}
}

// GetExtraList ...
func (repo *ExtraRepository) GetExtraList(ctx context.Context, branchStoreID int) ([]extra.Extra, error) {
	extraArray := []extra.Extra{}

	branchStoreObject := menu.BranchStore{}

	errorMsg := repo.db.Model(&menu.BranchStore{}).Where("id = ?", branchStoreID).Find(&branchStoreObject).Error
	if errorMsg != nil {
		return []extra.Extra{}, errorMsg
	}

	// search eveything with store id
	err := repo.db.Model(&extra.Extra{}).Where("store_id = ?", branchStoreObject.StoreGroupID).Find(&extraArray).Error
	if err != nil {
		return []extra.Extra{}, err
	}

	return extraArray, nil
}
