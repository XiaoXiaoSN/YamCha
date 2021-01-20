package order

import (
	"context"
	"yamcha/pkg/model"
)

// Service is a Order service
type Service interface {
	GetOrder(ctx context.Context, orderID int) (model.Order, error)
	GetGroupOrder(groupID string) (model.Order, error)
	OrderList(ctx context.Context, params model.OrderParams) ([]model.Order, error)
	CreateOrder(ctx context.Context, createOrderparams model.CreateOrderParams) (model.Order, error)
	UpdateOrder(ctx context.Context, createOrderparams model.CreateOrderParams) (model.Order, error)
	DeleteOrder(ctx context.Context, orderID int) error
	FinishOrder(groupID string) ([]model.PersonalOrder, error)
}
