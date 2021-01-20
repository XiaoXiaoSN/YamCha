package extra

import (
	"context"
	"yamcha/pkg/model"
)

// Service is a Order service
type Service interface {
	GetExtraList(ctx context.Context, storeID int) ([]model.Extra, error)
}
