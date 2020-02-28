package service

import (
	"context"
	"yamcha/pkg/api/menu"
)

// MenuService implement a menu service
type MenuService struct {
	MenuRepo menu.Repository
}

// NewMenuService make a menu servicer
func NewMenuService(menuRepo menu.Repository) menu.Service {
	return &MenuService{
		MenuRepo: menuRepo,
	}
}

// GetMenuList ...
func (svc *MenuService) GetMenuList(ctx context.Context, storeID int) ([]menu.Menu, error) {
	return svc.MenuRepo.GetMenuList(ctx, storeID)
}
