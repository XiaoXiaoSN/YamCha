package order

import (
	"context"
	"time"
)

// Order Object
type Order struct {
	ID        int       `gorm:"id" json:"id"`
	Creator   int       `gorm:"creator" json:"creator"`
	GroupID   int       `gorm:"group_id" json:"group_id"`
	Status    int       `gorm:"status" json:"status"`
	Price     string    `gorm:"price" json:"price"`
	Order     string    `gorm:"order" json:"order"`
	CreatedAt time.Time `gorm:"created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"updated_at" json:"updated_at"`
}

// Params from POST method
type Params struct {
	CreatorID     int `json:"creator_id" gorm:"creator_id" validate:"required"`
	GroupID       int `json:"group_id" gorm:"group_id" validate:"required"`
	BranchStoreID int `json:"branch_store_id" gorm:"branch_store_id" validate:"required"`
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
