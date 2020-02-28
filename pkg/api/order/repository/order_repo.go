package repository

import (
	"context"
	"encoding/json"
	"yamcha/pkg/api/order"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

// OrderRepository implement a order Repository
type OrderRepository struct {
	db *gorm.DB
}

// NewOrderRepository make a order Repository
func NewOrderRepository(db *gorm.DB) order.Repository {
	return &OrderRepository{
		db: db,
	}
}

// CreateOrder ...
func (repo *OrderRepository) CreateOrder(ctx context.Context, newOrder order.Order) (order.Order, error) {
	log.Println(newOrder)
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
	log.Println("GroupId", orderObject)
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
	err := repo.db.Model(&order.Order{}).Where("id = ? AND status = 1", orderID).Update("status", order.StatusOrderClose).Error
	if err != nil {
		return err
	}

	return nil
}

// UpdateOrder ...
func (repo *OrderRepository) UpdateOrder(ctx context.Context, newOrder order.Order) (order.Order, error) {
	log.Println(newOrder)
	err := repo.db.Model(&order.Order{}).Where("group_id = ? AND status = 1", newOrder.GroupID).Update("order", newOrder.Order).Error
	if err != nil {
		return order.Order{}, err
	}

	return newOrder, nil
}

// FinishOrder ...
func (repo *OrderRepository) FinishOrder(groupID string) ([]order.PersonalOrder, error) {
	orderList := order.Order{}
	err := repo.db.Model(&order.Order{}).Where("group_id = ? AND status = 1", groupID).Find(&orderList).Update("status", order.StatusOrderEnd).Error

	personalOrders := make([]order.PersonalOrder, 0)
	json.Unmarshal(orderList.Order, &personalOrders)

	log.Println("orderStruct")
	log.Println(personalOrders)

	orderList.OrderStruct = personalOrders

	if err != nil {
		return []order.PersonalOrder{}, err
	}

	return personalOrders, err
}
