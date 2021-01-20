package service

import (
	"context"
	"yamcha/pkg/service/menu"
	"yamcha/pkg/model"
	"yamcha/pkg/repository"
)

// MenuService implement a menu service
type MenuService struct {
	repo repository.Repository
}

// NewMenuService make a menu servicer
func NewMenuService(repo repository.Repository) menu.Service {
	return &MenuService{
		repo: repo,
	}
}

// GetMenuList ...
func (svc *MenuService) GetMenuList(ctx context.Context, storeID int) ([]model.Menu, error) {
	return svc.repo.GetMenuList(ctx, storeID)
}
