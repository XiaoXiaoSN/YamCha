package order

import (
	"context"
)

// Order Object
type Order struct {
	ID      int    `gorm:"id" json:"id"`
	Creator string `gorm:"name" json:"name"`
	Price   string `gorm:"price" json:"price"`
	Order   string `gorm:"order" json:"order"`
	// CreatedAt time.Time `gorm:"created_at" json:"created_at"`
	// UpdatedAt time.Time `gorm:"updated_at" json:"updated_at"`
}

// Service is a Order service
type Service interface {
	OrderList(ctx context.Context) ([]Order, error)
	CreateOrder(ctx context.Context, u Order) error
}

// Repository is a Order repo
type Repository interface {
	OrderList(ctx context.Context) ([]Order, error)
	CreateOrder(ctx context.Context, u Order) error
}
