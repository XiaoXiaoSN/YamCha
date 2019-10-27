package order

import (
	"context"
)

// Order Object
type Order struct {
	ID      int    `gorm:"id" json:"id"`
	Creator string `gorm:"creator" json:"creator"`
	Group   string `gorm:"group_id" json:"group_id"`
	Price   string `gorm:"price" json:"price"`
	Order   string `gorm:"order" json:"order"`
	// CreatedAt time.Time `gorm:"created_at" json:"created_at"`
	// UpdatedAt time.Time `gorm:"updated_at" json:"updated_at"`
}

type Params struct {
	Creator string `gorm:"creator" json:"creator"`
	Group   string `gorm:"group_id" json:"group_id"`
}

// Service is a Order service
type Service interface {
	OrderList(ctx context.Context, id string) (Order, error)
	CreateOrder(ctx context.Context, param Params) (Order, error)
}

// Repository is a Order repo
type Repository interface {
	OrderList(ctx context.Context, id string) (Order, error)
	CreateOrder(ctx context.Context, param Params) (Order, error)
}
