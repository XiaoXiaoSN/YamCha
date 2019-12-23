package repository

import (
	"context"
	"yamcha/pkg/api/order"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
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
func (repo *OrderRepository) CreateOrder(ctx context.Context, newOrder order.Order) (order.Order, error) {
	newOrder.Order = []order.PersonalOrder{}
	err := repo.db.Model(&order.Order{}).Create(&newOrder).Error
	if err != nil {
		return order.Order{}, err
	}

	return newOrder, nil
}

// GetOrder ...
func (repo *OrderRepository) GetOrder(ctx context.Context, orderID int) (order.Order, error) {
	orderObject := order.Order{}

	err := repo.db.Model(&order.Order{}).Where("id = ?", orderID).First(&orderObject).Error
	if err != nil {
		return order.Order{}, err
	}

	return orderObject, nil
}

// GetGroupOrder ...
func (repo *OrderRepository) GetGroupOrder(groupID string) (order.Order, error) {
	orderObject := order.Order{}
	log.Println("GroupId", groupID)
	err := repo.db.Model(&order.Order{}).Where("group_id = ? AND status = 1", groupID).First(&orderObject).Error
	if err != nil {
		return order.Order{}, err
	}

	return orderObject, nil
}

// OrderList ...
func (repo *OrderRepository) OrderList(ctx context.Context, params order.Params) ([]order.Order, error) {
	orderList := []order.Order{}

	model := repo.db.Model(&order.Order{})

	// inject the filters
	if params.CreatorID != nil {
		model = model.Where("creator_id = ?", params.CreatorID)
	}
	if params.GroupID != nil {
		model = model.Where("group_id = ?", params.GroupID)
	}

	err := model.Find(&orderList).Error
	if err != nil {
		return []order.Order{}, err
	}

	return orderList, nil
}

// DeleteOrder ...
func (repo *OrderRepository) DeleteOrder(ctx context.Context, orderID int) error {
	// orderObject := order.Order{}
	err := repo.db.Model(&order.Order{}).Where("id = ? AND status = 1", orderID).Update("status", 2).Error
	if err != nil {
		return err
	}

	return nil
}

// UpdateOrder ...
func (repo *OrderRepository) UpdateOrder(ctx context.Context, newOrder order.Order) (order.Order, error) {
	log.Println(newOrder)
	err := repo.db.Model(&order.Order{}).Where("id = ? AND status = 1", newOrder.ID).Update("order", newOrder.Order).Error
	if err != nil {
		return order.Order{}, err
	}

	return newOrder, nil
}
