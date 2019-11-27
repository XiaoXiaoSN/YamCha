package order

import (
	"context"
	"encoding/json"
	"time"
)

const (
	// StatusOrderOpen is that the order keep open
	StatusOrderOpen int8 = 1
	// StatusOrderClose is that the order be closed by creator
	StatusOrderClose int8 = 2
	// StatusOrderEnd is that the order end successfully
	StatusOrderEnd int8 = 3
)

// Order Object
type Order struct {
	ID        int             `gorm:"id" json:"id"`
	Creator   string          `gorm:"creator" json:"creator"`
	GroupID   string          `gorm:"group_id" json:"group_id"`
	Status    int8            `gorm:"status" json:"status"`
	Price     int             `gorm:"price" json:"price"`
	Order     json.RawMessage `gorm:"order" json:"order"`
	CreatedAt time.Time       `gorm:"created_at" json:"created_at"`
	UpdatedAt time.Time       `gorm:"updated_at" json:"updated_at"`
}

// Params for filter order list
type Params struct {
	CreatorID *int    `json:"creator_id" query:"creator_id" form:"creator_id"`
	GroupID   *string `json:"group_id" query:"group_id" form:"group_id"`
}

// CreateOrderParams for create a new order
type CreateOrderParams struct {
	CreatorID     string          `json:"creator_id" validate:"required"`
	GroupID       string          `json:"group_id" validate:"required"`
	BranchStoreID int             `json:"branch_store_id" validate:"required"`
	Order         json.RawMessage `gorm:"order" json:"order"`
}

// Service is a Order service
type Service interface {
	GetOrder(ctx context.Context, orderID int) (Order, error)
	GetGroupOrder(groupID string) (Order, error)
	OrderList(ctx context.Context, params Params) ([]Order, error)
	CreateOrder(ctx context.Context, createOrderparams CreateOrderParams) (Order, error)
}

// Repository is a Order repo
type Repository interface {
	GetOrder(ctx context.Context, orderID int) (Order, error)
	GetGroupOrder(groupID string) (Order, error)
	OrderList(ctx context.Context, params Params) ([]Order, error)
	CreateOrder(ctx context.Context, order Order) (Order, error)
}
