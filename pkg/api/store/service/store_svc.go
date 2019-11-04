package service

import (
	"context"
	"yamcha/pkg/api/store"
)

// StoreService implment a user service
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
func (svc *StoreService) CreateStore(ctx context.Context, u store.Store) (store.Store, error) {
	return svc.StoreRepo.CreateStore(ctx, u)
}

// StoreList ...
func (svc *StoreService) StoreList(ctx context.Context) ([]store.Store, error) {
	return svc.StoreRepo.StoreList(ctx)
}

// BranchStoreList ...
func (svc *StoreService) BranchStoreList(ctx context.Context, storeID int) ([]store.BranchStore, error) {
	return svc.StoreRepo.BranchStoreList(ctx, storeID)
}
