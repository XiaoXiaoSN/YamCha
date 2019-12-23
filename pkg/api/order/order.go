package order

import (
	"context"
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
	ID            int             `gorm:"id" json:"id"`
	Creator       string          `gorm:"creator_id" json:"creator"`
	GroupID       string          `gorm:"group_id" json:"group_id"`
	Status        int8            `gorm:"status" json:"status"`
	Price         int             `gorm:"price" json:"price"`
	BranchStoreID int             `gorm:"branch_store_id" json:"branch_store_id"`
	Order         []PersonalOrder `gorm:"order" json:"order"`
	CreatedAt     time.Time       `gorm:"created_at" json:"created_at"`
	UpdatedAt     time.Time       `gorm:"updated_at" json:"updated_at"`
}

// Params for filter order list
type Params struct {
	CreatorID *int    `json:"creator_id" query:"creator_id" form:"creator_id"`
	GroupID   *string `json:"group_id" query:"group_id" form:"group_id"`
}

// PersonalOrder ...
type PersonalOrder struct {
	UserID    string `gorm:"user" form:"user" json:"user"`
	ProdustID string `gorm:"product" form:"product" json:"product"`
	Size      string `gorm:"size" form:"size" json:"size"`
	Sweet     string `gorm:"sweet" form:"sweet" json:"sweet"`
	Ice       string `gorm:"ice" form:"ice" json:"ice"`
	Price     string `gorm:"price" form:"price" json:"price"`
}

// CreateOrderParams for create a new order
type CreateOrderParams struct {
	CreatorID     string          `gorm:"creator_id" form:"creator_id" json:"creator_id"`
	GroupID       string          `gorm:"group_id" form:"group_id" json:"group_id"`
	BranchStoreID int             `gorm:"branch_store_id" form:"branch_store_id" json:"branch_store_id"`
	Order         []PersonalOrder `gorm:"order" json:"order"`
}

// Service is a Order service
type Service interface {
	GetOrder(ctx context.Context, orderID int) (Order, error)
	GetGroupOrder(groupID string) (Order, error)
	OrderList(ctx context.Context, params Params) ([]Order, error)
	CreateOrder(ctx context.Context, createOrderparams CreateOrderParams) (Order, error)
	UpdateOrder(ctx context.Context, createOrderparams CreateOrderParams) (Order, error)
	DeleteOrder(ctx context.Context, orderID int) error
}

// Repository is a Order repo
type Repository interface {
	GetOrder(ctx context.Context, orderID int) (Order, error)
	GetGroupOrder(groupID string) (Order, error)
	OrderList(ctx context.Context, params Params) ([]Order, error)
	CreateOrder(ctx context.Context, order Order) (Order, error)
	UpdateOrder(ctx context.Context, order Order) (Order, error)
	DeleteOrder(ctx context.Context, orderID int) error
}
