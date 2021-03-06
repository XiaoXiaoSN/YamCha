package db

import (
	"context"
	"encoding/json"

	"yamcha/pkg/model"
)

// CreateOrder ...
func (repo *dbRepository) CreateOrder(ctx context.Context, newOrder *model.Order) error {
	err := repo.db.Model(&model.Order{}).Create(newOrder).Error
	if err != nil {
		return err
	}
	return nil
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
func (repo *dbRepository) GetGroupOrder(ctx context.Context, groupID string) (model.Order, error) {
	var order model.Order

	err := repo.db.Model(&model.Order{}).
		Where("group_id = ?", groupID).
		Where("status = ?", model.OrderStatusOpen).
		First(&order).Error
	if err != nil {
		return model.Order{}, err
	}

	return order, nil
}

// OrderList ...
func (repo *dbRepository) OrderList(ctx context.Context, params model.OrderParams) ([]model.Order, error) {
	orderList := []model.Order{}

	model := repo.db.Model(&model.Order{})

	// parse filters
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
		Where("id = ?", orderID).
		Where("status = ?", model.OrderStatusOpen).
		Update("status", model.OrderStatusClose).Error
	if err != nil {
		return err
	}

	return nil
}

// UpdateOrder ...
func (repo *dbRepository) UpdateOrder(ctx context.Context, order model.Order) (model.Order, error) {
	err := repo.db.Model(&model.Order{}).
		Where("group_id = ?", order.GroupID).
		Where("status = ?", model.OrderStatusOpen).
		Update("order", order.Order).Error
	if err != nil {
		return model.Order{}, err
	}

	return order, nil
}

// FinishOrder ...
func (repo *dbRepository) FinishOrder(ctx context.Context, groupID string) ([]model.PersonalOrder, error) {
	order := model.Order{}
	err := repo.db.Model(&model.Order{}).
		Where("group_id = ?", groupID).
		Where("status = ?", model.OrderStatusOpen).
		Find(&order).
		Update("status", model.OrderStatusEnd).Error

	personalOrders := make([]model.PersonalOrder, 0)
	err = json.Unmarshal(order.Order, &personalOrders)
	if err != nil {
		return []model.PersonalOrder{}, err
	}

	return personalOrders, err
}
