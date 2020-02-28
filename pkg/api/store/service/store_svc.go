package service

import (
	"context"
	"yamcha/pkg/api/store"
)

// StoreService implement a user service
type StoreService struct {
	StoreRepo store.Repository
}

// NewStoreService make a user servicer
func NewStoreService(storeRepo store.Repository) store.Service {
	return &StoreService{
		StoreRepo: storeRepo,
	}
}

// CreateStore ...
func (svc *StoreService) CreateStore(ctx context.Context, targetStore store.Store) (store.Store, error) {
	return svc.StoreRepo.CreateStore(ctx, targetStore)
}

// GetStore ...
func (svc *StoreService) GetStore(ctx context.Context, storeID int) (store.Store, error) {
	return svc.StoreRepo.GetStore(ctx, storeID)
}

// StoreList ...
func (svc *StoreService) StoreList(ctx context.Context) ([]store.Store, error) {
	return svc.StoreRepo.StoreList(ctx)
}

// BranchStoreList ...
func (svc *StoreService) BranchStoreList(ctx context.Context, storeID int) ([]store.BranchStore, error) {
	return svc.StoreRepo.BranchStoreList(ctx, storeID)
}

// CreateBranchStore ...
func (svc *StoreService) CreateBranchStore(ctx context.Context, branchStore store.BranchStore) (store.BranchStore, error) {
	return svc.StoreRepo.CreateBranchStore(ctx, branchStore)

}
