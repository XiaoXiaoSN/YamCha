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
func (svc *OrderService) CreateOrder(ctx context.Context, u order.Order) error {
	return nil
}

// OrderList ...
func (svc *OrderService) OrderList(ctx context.Context) ([]order.Order, error) {
	return svc.OrderRepo.OrderList(ctx)
}
