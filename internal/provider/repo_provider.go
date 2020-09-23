package provider

import (
	extraRepo "yamcha/pkg/api/extra/repository"
	menuRepo "yamcha/pkg/api/menu/repository"
	orderRepo "yamcha/pkg/api/order/repository"
	storeRepo "yamcha/pkg/api/store/repository"
	userRepo "yamcha/pkg/api/user/repository"

	"github.com/google/wire"
)

// RepoSet ...
var RepoSet = wire.NewSet(
	extraRepo.NewExtraRepository,
	menuRepo.NewMenuRepository,
	orderRepo.NewOrderRepository,
	storeRepo.NewStoreRepository,
	userRepo.NewUserRepository,
)
