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
func (repo *OrderRepository) CreateOrder(ctx context.Context, param order.Params) (order.Order, error) {
	orderObject := order.Order{
		Group:   param.Group,
		Creator: param.Creator,
	}
	err := repo.db.Model(&order.Order{}).Create(&orderObject).Error
	if err != nil {
		return order.Order{}, err
	}

	return orderObject, nil

}

// OrderList ...
func (repo *OrderRepository) OrderList(ctx context.Context, id string) (order.Order, error) {
	orderObject := order.Order{}

	err := repo.db.Model(&order.Order{}).Where("id = ?", id).First(&orderObject).Error
	if err != nil {
		return order.Order{}, err
	}

	return orderObject, nil
}
