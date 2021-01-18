package service

import (
	"context"
	"yamcha/pkg/api/store"
	"yamcha/pkg/model"
	"yamcha/pkg/repository"
)

// StoreService implement a user service
type StoreService struct {
	repo repository.Repository
}

// NewStoreService make a user servicer
func NewStoreService(repo repository.Repository) store.Service {
	return &StoreService{
		repo: repo,
	}
}

// CreateStore ...
func (svc *StoreService) CreateStore(ctx context.Context, targetStore model.Store) (model.Store, error) {
	return svc.repo.CreateStore(ctx, targetStore)
}

// GetStore ...
func (svc *StoreService) GetStore(ctx context.Context, storeID int) (model.Store, error) {
	return svc.repo.GetStore(ctx, storeID)
}

// StoreList ...
func (svc *StoreService) StoreList(ctx context.Context) ([]model.Store, error) {
	return svc.repo.StoreList(ctx)
}

// BranchStoreList ...
func (svc *StoreService) BranchStoreList(ctx context.Context, storeID int) ([]model.BranchStore, error) {
	return svc.repo.BranchStoreList(ctx, storeID)
}

// CreateBranchStore ...
func (svc *StoreService) CreateBranchStore(ctx context.Context, branchStore model.BranchStore) (model.BranchStore, error) {
	return svc.repo.CreateBranchStore(ctx, branchStore)

}
