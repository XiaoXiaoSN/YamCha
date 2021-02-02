package store

import (
	"context"
	"yamcha/pkg/model"
)

// Service is a store service
type Service interface {
	GetStore(ctx context.Context, storeID int) (model.Store, error)
	StoreList(ctx context.Context) ([]model.Store, error)
	CreateStore(ctx context.Context, s *model.Store) error

	// branch store
	BranchStoreList(ctx context.Context, storeID int) ([]model.BranchStore, error)
	CreateBranchStore(ctx context.Context, branchStore *model.BranchStore) error
}
