package service

import (
	"context"
	"yamcha/pkg/api/store"
)

// UserService implment a user service
type StoreService struct {
	StoreRepo store.Repository
}

// NewUserService make a user servicer
func NewStoreService(storeRepo store.Repository) store.Service {
	return &StoreService{
		StoreRepo: storeRepo,
	}
}

// CreateUser ...
func (svc *StoreService) CreateStore(ctx context.Context, u store.Store) error {
	return nil
}

// UserList ...
func (svc *StoreService) StoreList(ctx context.Context) ([]store.Store, error) {
	return svc.StoreRepo.StoreList(ctx)
}
