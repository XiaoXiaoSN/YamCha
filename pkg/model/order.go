package model

import (
	"encoding/json"
	"time"
)

// OrderStatus ..
type OrderStatus int8

const (
	// OrderStatusOpen is that the order keep open
	OrderStatusOpen OrderStatus = 1
	// OrderStatusClose is that the order be closed by creator
	OrderStatusClose OrderStatus = 2
	// OrderStatusEnd is that the order end successfully
	OrderStatusEnd OrderStatus = 3
)

// Order Object
type Order struct {
	ID            int             `gorm:"id" json:"id"`
	Creator       string          `gorm:"creator_id" json:"creator"`
	GroupID       string          `gorm:"group_id" json:"group_id"`
	Status        OrderStatus     `gorm:"status" json:"status"`
	Price         int             `gorm:"price" json:"price"`
	BranchStoreID int             `gorm:"branch_store_id" json:"branch_store_id"`
	Order         json.RawMessage `gorm:"order" json:"order,omitempty"`
	CreatedAt     time.Time       `gorm:"created_at" json:"created_at"`
	UpdatedAt     time.Time       `gorm:"updated_at" json:"updated_at"`
	OrderStruct   []PersonalOrder
}

// OrderParams for filter order list
type OrderParams struct {
	CreatorID *string `json:"creator_id" query:"creator_id" form:"creator_id"`
	GroupID   *string `json:"group_id" query:"group_id" form:"group_id"`
}

// PersonalOrder ...
type PersonalOrder struct {
	UserID    string `gorm:"user" form:"user" json:"user"`
	ProductID string `gorm:"product" form:"product" json:"product"`
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
