package extra

import (
	"context"
)

// Extra Object
type Extra struct {
	ID      int    `gorm:"id" json:"id"`
	Name    string `gorm:"name" json:"name"`
	StoreID string `gorm:"store_id" json:"store_id"`
	Price   int8   `gorm:"price" json:"price"`
}

// Service is a Order service
type Service interface {
	GetExtraList(ctx context.Context, storeID int) ([]Extra, error)
}

// Repository is a Order repo
type Repository interface {
	GetExtraList(ctx context.Context, storeID int) ([]Extra, error)
}
