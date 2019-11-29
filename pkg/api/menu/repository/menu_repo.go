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
func (repo *MenuRepository) GetMenuList(ctx context.Context, storeID int) ([]menu.Menu, error) {
	menuObject := []menu.Menu{}
	// search eveything with store id
	err := repo.db.Model(&menu.Menu{}).Where("store_id = ?", storeID).Find(&menuObject).Error
	if err != nil {
		return []menu.Menu{}, err
	}

	return menuObject, nil
}
