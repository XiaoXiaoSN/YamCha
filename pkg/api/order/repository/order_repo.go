package repository

import (
	"context"
	"yamcha/pkg/api/order"

	"github.com/jinzhu/gorm"
)

// OrderRepository implment a order Repository
type OrderRepository struct {
	db *gorm.DB
}

// NewOrderRepository make a order Repositoryr
func NewOrderRepository(db *gorm.DB) order.Repository {
	return &OrderRepository{
		db: db,
	}
}

// CreateOrder ...
func (repo *OrderRepository) CreateOrder(ctx context.Context, u order.Order) error {
	return nil
}

// OrderList ...
func (repo *OrderRepository) OrderList(ctx context.Context) ([]order.Order, error) {
	orderList := []order.Order{}

	err := repo.db.Model(&order.Order{}).Find(&orderList).Error
	if err != nil {
		return []order.Order{}, err
	}

	return orderList, nil
}
