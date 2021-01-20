package menu

import (
	"context"
	"yamcha/pkg/model"
)

// Service is a Order service
type Service interface {
	GetMenuList(ctx context.Context, storeID int) ([]model.Menu, error)
}
