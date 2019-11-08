package service

import (
	"context"
	"yamcha/pkg/api/order"
)

// OrderService implment a order service
type OrderService struct {
	OrderRepo order.Repository
}

// NewOrderService make a order servicer
func NewOrderService(orderRepo order.Repository) order.Service {
	return &OrderService{
		OrderRepo: orderRepo,
	}
}

// CreateOrder ...
func (svc *OrderService) CreateOrder(ctx context.Context, cParams order.CreateOrderParams) (order.Order, error) {
	orderObject := order.Order{
		GroupID: cParams.GroupID,
		Creator: cParams.CreatorID,
		Price:   0,
		Status:  order.StatusOrderOpen,
	}

	return svc.OrderRepo.CreateOrder(ctx, orderObject)
}

// GetOrder ...
func (svc *OrderService) GetOrder(ctx context.Context, orderID int) (order.Order, error) {
	return svc.OrderRepo.GetOrder(ctx, orderID)
}

// OrderList ...
func (svc *OrderService) OrderList(ctx context.Context, params order.Params) ([]order.Order, error) {
	return svc.OrderRepo.OrderList(ctx, params)
}
