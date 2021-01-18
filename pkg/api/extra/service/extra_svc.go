package service

import (
	"context"
	"yamcha/pkg/api/extra"
	"yamcha/pkg/model"
	"yamcha/pkg/repository"
)

// ExtraService implement a extra service
type ExtraService struct {
	repo repository.Repository
}

// NewExtraService make a extra servicer
func NewExtraService(repo repository.Repository) extra.Service {
	return &ExtraService{
		repo: repo,
	}
}

// GetExtraList ...
func (svc *ExtraService) GetExtraList(ctx context.Context, storeID int) ([]model.Extra, error) {
	return svc.repo.GetExtraList(ctx, storeID)
}
