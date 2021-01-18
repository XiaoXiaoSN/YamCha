package service

import (
	"context"
	"encoding/json"
	"yamcha/pkg/api/order"
	"yamcha/pkg/model"
	"yamcha/pkg/repository"

	log "github.com/sirupsen/logrus"
)

// OrderService implement a order service
type OrderService struct {
	repo repository.Repository
}

// NewOrderService make a order servicer
func NewOrderService(repo repository.Repository) order.Service {
	return &OrderService{
		repo: repo,
	}
}

// CreateOrder ...
func (svc *OrderService) CreateOrder(ctx context.Context, cParams model.CreateOrderParams) (model.Order, error) {
	orderObject := model.Order{
		GroupID:       cParams.GroupID,
		Creator:       cParams.CreatorID,
		BranchStoreID: cParams.BranchStoreID,
		Price:         0,
		Order:         []byte("{}"),
		Status:        model.OrderStatusOpen,
	}

	return svc.repo.CreateOrder(ctx, orderObject)
}

// GetGroupOrder ...
func (svc *OrderService) GetGroupOrder(groupID string) (model.Order, error) {
	log.Println("in func")
	return svc.repo.GetGroupOrder(groupID)
}

// GetOrder ...
func (svc *OrderService) GetOrder(ctx context.Context, orderID int) (model.Order, error) {
	return svc.repo.GetOrder(ctx, orderID)
}

// OrderList ...
func (svc *OrderService) OrderList(ctx context.Context, params model.OrderParams) ([]model.Order, error) {
	return svc.repo.OrderList(ctx, params)
}

// DeleteOrder ...
func (svc *OrderService) DeleteOrder(ctx context.Context, orderID int) error {
	return svc.repo.DeleteOrder(ctx, orderID)
}

// UpdateOrder ...
func (svc *OrderService) UpdateOrder(ctx context.Context, cParams model.CreateOrderParams) (model.Order, error) {
	stringJSON, _ := json.Marshal(cParams.Order)
	orderObject := model.Order{
		GroupID: cParams.GroupID,
		Creator: cParams.CreatorID,
		Price:   0,
		Status:  model.OrderStatusOpen,
		Order:   []byte(stringJSON),
	}
	return svc.repo.UpdateOrder(ctx, orderObject)
}

// FinishOrder ...
func (svc *OrderService) FinishOrder(groupID string) ([]model.PersonalOrder, error) {
	return svc.repo.FinishOrder(groupID)
}
