package repository

import (
	// extraRepo "yamcha/pkg/api/extra/repository"
	// menuRepo "yamcha/pkg/api/menu/repository"
	// orderRepo "yamcha/pkg/api/order/repository"
	// storeRepo "yamcha/pkg/api/store/repository"
	// userRepo "yamcha/pkg/api/user/repository"

	"yamcha/pkg/api/extra"
	"yamcha/pkg/api/menu"
	"yamcha/pkg/api/order"
	"yamcha/pkg/api/store"
	"yamcha/pkg/api/user"
)

// Repository ...
type Repository interface {
	// extra.Repository
	// menu.Repository
	// order.Repository
	// store.Repository
	// user.Repository
}

type _repository struct {
	// extraRepo.ExtraRepository
	// menuRepo.MenuRepository
	// orderRepo.OrderRepository
	// storeRepo.StoreRepository
	// userRepo.UserRepository
}

// NewRepo ...
func NewRepo(
	extraRepo extra.Repository,
	menuRepo menu.Repository,
	orderRepo order.Repository,
	storeRepo store.Repository,
	userRepo user.Repository,
) Repository {
	return &_repository{
		// extraRepo,
		// menuRepo,
		// orderRepo,
		// storeRepo,
		// userRepo,
	}
}
