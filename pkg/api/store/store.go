package store

import (
	"context"
)

// Store Object
type Store struct {
	ID        int    `gorm:"id" json:"id"`
	GroupName string `gorm:"group_name" json:"group_name"`
	// CreatedAt time.Time `gorm:"created_at" json:"created_at"`
	// UpdatedAt time.Time `gorm:"updated_at" json:"updated_at"`
}

// BranchStore Object
type BranchStore struct {
	ID           int    `gorm:"id" json:"id"`
	Name         string `gorm:"name" json:"name"`
	StoreGroupID int    `gorm:"store_group_id" json:"store_group_id"`
	Address      string `gorm:"address" json:"address"`
	Phone        string `gorm:"phone" json:"phone"`
	// CreatedAt time.Time `gorm:"created_at" json:"created_at"`
	// UpdatedAt time.Time `gorm:"updated_at" json:"updated_at"`
}

// TableName of Store
func (s *Store) TableName() string {
	return "store_groups"
}

// TableName of BranchStore
func (s *BranchStore) TableName() string {
	return "stores"
}

// Service is a store service
type Service interface {
	StoreList(ctx context.Context) ([]Store, error)
	BranchStoreList(ctx context.Context, id string) ([]BranchStore, error)
	CreateStore(ctx context.Context, u Store) (Store, error)
}

// Repository is a store repo
type Repository interface {
	StoreList(ctx context.Context) ([]Store, error)
	BranchStoreList(ctx context.Context, id string) ([]BranchStore, error)
	CreateStore(ctx context.Context, u Store) (Store, error)
}
