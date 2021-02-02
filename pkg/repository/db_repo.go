package repository

import (
	"context"
	"yamcha/pkg/model"
)

// Repository ...
type Repository interface {
	// user
	UserList(ctx context.Context) ([]model.User, error)
	CreateUser(ctx context.Context, u *model.User) error

	// store
	GetStore(ctx context.Context, storeID int) (model.Store, error)
	StoreList(ctx context.Context) ([]model.Store, error)
	CreateStore(ctx context.Context, s *model.Store) error

	// branch store
	BranchStoreList(ctx context.Context, storeID int) ([]model.BranchStore, error)
	CreateBranchStore(ctx context.Context, branchStore *model.BranchStore) error

	// order
	GetOrder(ctx context.Context, orderID int) (model.Order, error)
	GetGroupOrder(groupID string) (model.Order, error)
	OrderList(ctx context.Context, params model.OrderParams) ([]model.Order, error)
	CreateOrder(ctx context.Context, order *model.Order) error
	UpdateOrder(ctx context.Context, order model.Order) (model.Order, error)
	DeleteOrder(ctx context.Context, orderID int) error
	FinishOrder(groupID string) ([]model.PersonalOrder, error)

	// menu
	GetMenuList(ctx context.Context, storeID int) ([]model.Menu, error)

	// extra
	GetExtraList(ctx context.Context, storeID int) ([]model.Extra, error)
}
