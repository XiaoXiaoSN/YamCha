package service

import (
	"context"
	"yamcha/pkg/api/extra"
)

// ExtraService implement a extra service
type ExtraService struct {
	ExtraRepo extra.Repository
}

// NewExtraService make a extra servicer
func NewExtraService(extraRepo extra.Repository) extra.Service {
	return &ExtraService{
		ExtraRepo: extraRepo,
	}
}

// GetExtraList ...
func (svc *ExtraService) GetExtraList(ctx context.Context, storeID int) ([]extra.Extra, error) {
	return svc.ExtraRepo.GetExtraList(ctx, storeID)
}
