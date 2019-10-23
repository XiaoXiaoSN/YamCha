package store

import (
	"context"
)

// Store Object
type Store struct {
	ID      int    `gorm:"id" json:"id"`
	Name    string `gorm:"name" json:"name"`
	Address string `gorm:"address" json:"address"`
	Phone   string `gorm:"phone" json:"phone"`
	// CreatedAt time.Time `gorm:"created_at" json:"created_at"`
	// UpdatedAt time.Time `gorm:"updated_at" json:"updated_at"`
}

// Service is a store service
type Service interface {
	StoreList(ctx context.Context) ([]Store, error)
	CreateStore(ctx context.Context, u Store) error
}

// Repository is a store repo
type Repository interface {
	StoreList(ctx context.Context) ([]Store, error)
	CreateStore(ctx context.Context, u Store) error
}
