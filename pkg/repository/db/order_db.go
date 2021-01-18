package db

import (
	"context"
	"encoding/json"
	"log"

	"yamcha/pkg/model"
)

// CreateOrder ...
func (repo *dbRepository) CreateOrder(ctx context.Context, newOrder model.Order) (model.Order, error) {
	log.Println(newOrder)
	err := repo.db.Model(&model.Order{}).Create(&newOrder).Error
	if err != nil {
		return model.Order{}, err
	}

	return newOrder, nil
}

// GetOrder ...
func (repo *dbRepository) GetOrder(ctx context.Context, orderID int) (model.Order, error) {
	orderObject := model.Order{}

	err := repo.db.Model(&model.Order{}).
		Where("id = ?", orderID).
		First(&orderObject).Error
	if err != nil {
		return model.Order{}, err
	}

	return orderObject, nil
}

// GetGroupOrder ...
func (repo *dbRepository) GetGroupOrder(groupID string) (model.Order, error) {
	orderObject := model.Order{}
	log.Println("GroupId", groupID)
	err := repo.db.Model(&model.Order{}).
		Where("group_id = ? AND status = 1", groupID).
		First(&orderObject).Error
	log.Printf("GroupId %+v", orderObject)
	if err != nil {
		return model.Order{}, err
	}

	return orderObject, nil
}

// OrderList ...
func (repo *dbRepository) OrderList(ctx context.Context, params model.OrderParams) ([]model.Order, error) {
	orderList := []model.Order{}

	model := repo.db.Model(&model.Order{})

	// inject the filters
	if params.CreatorID != nil {
		model = model.Where("creator_id = ?", params.CreatorID)
	}
	if params.GroupID != nil {
		model = model.Where("group_id = ?", params.GroupID)
	}

	err := model.Find(&orderList).Error
	if err != nil {
		return nil, err
	}

	return orderList, nil
}

// DeleteOrder ...
func (repo *dbRepository) DeleteOrder(ctx context.Context, orderID int) error {
	// orderObject := model.Order{}
	err := repo.db.Model(&model.Order{}).
		Where("id = ? AND status = 1", orderID).
		Update("status", model.OrderStatusClose).Error
	if err != nil {
		return err
	}

	return nil
}

// UpdateOrder ...
func (repo *dbRepository) UpdateOrder(ctx context.Context, newOrder model.Order) (model.Order, error) {
	err := repo.db.Model(&model.Order{}).
		Where("group_id = ? AND status = 1", newOrder.GroupID).
		Update("order", newOrder.Order).Error
	if err != nil {
		return model.Order{}, err
	}

	return newOrder, nil
}

// FinishOrder ...
func (repo *dbRepository) FinishOrder(groupID string) ([]model.PersonalOrder, error) {
	orderList := model.Order{}
	err := repo.db.Model(&model.Order{}).
		Where("group_id = ? AND status = 1", groupID).
		Find(&orderList).
		Update("status", model.OrderStatusEnd).Error

	personalOrders := make([]model.PersonalOrder, 0)
	json.Unmarshal(orderList.Order, &personalOrders)

	orderList.OrderStruct = personalOrders

	if err != nil {
		return []model.PersonalOrder{}, err
	}

	return personalOrders, err
}
