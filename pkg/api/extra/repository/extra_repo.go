package repository

import (
	"context"
	"yamcha/pkg/api/extra"

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
func (repo *ExtraRepository) GetExtraList(ctx context.Context, storeID int) ([]extra.Extra, error) {
	extraObject := []extra.Extra{}
	// search eveything with store id
	err := repo.db.Model(&extra.Extra{}).Where("store_id = ?", storeID).Find(&extraObject).Error
	if err != nil {
		return []extra.Extra{}, err
	}

	return extraObject, nil
}
