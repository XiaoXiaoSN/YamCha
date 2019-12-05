package repository

import (
	"context"
	"yamcha/pkg/api/menu"

	"github.com/jinzhu/gorm"
)

// MenuRepository implment a menu Repository
type MenuRepository struct {
	db *gorm.DB
}

// NewMenuRepository make a menu Repositoryr
func NewMenuRepository(db *gorm.DB) menu.Repository {
	return &MenuRepository{
		db: db,
	}
}

// GetMenuList ...
func (repo *MenuRepository) GetMenuList(ctx context.Context, branchStoreID int) ([]menu.Menu, error) {
	branchStoreObject := menu.BranchStore{}
	menuObject := []menu.Menu{}

	errorMsg := repo.db.Model(&menu.BranchStore{}).Where("id = ?", branchStoreID).Find(&branchStoreObject).Error
	if errorMsg != nil {
		return []menu.Menu{}, errorMsg
	}
	// log.Println(branchStoreObject.StoreID)
	// search eveything with store id
	err := repo.db.Model(&menu.Menu{}).Where("store_id = ?", branchStoreObject.StoreGroupID).Find(&menuObject).Error
	if err != nil {
		return []menu.Menu{}, err
	}

	return menuObject, nil

}
