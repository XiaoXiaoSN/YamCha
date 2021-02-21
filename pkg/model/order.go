package model

import (
	"time"
	"yamcha/internal/gormutil"
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
// TODO: 也許有些地方需要加一些 uniq key
type Order struct {
	ID            int             `gorm:"column:id" json:"id"`
	CreatorID     string          `gorm:"column:creator_id" json:"creator"`
	GroupID       string          `gorm:"column:group_id" json:"group_id"`
	Status        OrderStatus     `gorm:"column:status" json:"status"`
	Price         int             `gorm:"column:price" json:"price"`
	BranchStoreID int             `gorm:"column:branch_store_id" json:"branch_store_id"`
	Order         gormutil.JSON   `gorm:"column:order" json:"order,omitempty"`
	CreatedAt     time.Time       `gorm:"column:created_at" json:"created_at"`
	UpdatedAt     time.Time       `gorm:"column:updated_at" json:"updated_at"`
	OrderStruct   []PersonalOrder `gorm:"-"`
}

// OrderParams for filter order list
type OrderParams struct {
	CreatorID *string `json:"creator_id" query:"creator_id" form:"creator_id"`
	GroupID   *string `json:"group_id" query:"group_id" form:"group_id"`
}

// PersonalOrder ...
type PersonalOrder struct {
	UserID    string `gorm:"column:user" form:"user" json:"user"`
	ProductID string `gorm:"column:product" form:"product" json:"product"`
	Size      string `gorm:"column:size" form:"size" json:"size"`
	Sweet     string `gorm:"column:sweet" form:"sweet" json:"sweet"`
	Ice       string `gorm:"column:ice" form:"ice" json:"ice"`
	Price     string `gorm:"column:price" form:"price" json:"price"`
}

// CreateOrderParams for create a new order
type CreateOrderParams struct {
	CreatorID     string          `gorm:"column:creator_id" form:"creator_id" json:"creator_id"`
	GroupID       string          `gorm:"column:group_id" form:"group_id" json:"group_id"`
	BranchStoreID int             `gorm:"column:branch_store_id" form:"branch_store_id" json:"branch_store_id"`
	Order         []PersonalOrder `gorm:"column:order" json:"order"`
}
