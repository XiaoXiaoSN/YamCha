package menu

import (
	"context"
	"time"
)

// Menu Object
type Menu struct {
	ID        int       `gorm:"id" json:"id"`
	Name      string    `gorm:"name" json:"name"`
	StoreID   string    `gorm:"store_id" json:"store_id"`
	Price     int8      `gorm:"price" json:"price"`
	Size      int       `gorm:"size" json:"size"`
	CreatedAt time.Time `gorm:"created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"updated_at" json:"updated_at"`
}

// Service is a Order service
type Service interface {
	GetMenuList(ctx context.Context, storeID int) ([]Menu, error)
}

// Repository is a Order repo
type Repository interface {
	GetMenuList(ctx context.Context, storeID int) ([]Menu, error)
}
