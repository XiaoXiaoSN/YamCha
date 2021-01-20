package provider

import (
	extraSvc "yamcha/pkg/service/extra/service"
	menuSvc "yamcha/pkg/service/menu/service"
	orderSvc "yamcha/pkg/service/order/service"
	storeSvc "yamcha/pkg/service/store/service"
	userSvc "yamcha/pkg/service/user/service"

	"github.com/google/wire"
)

// ServiceSet ...
var ServiceSet = wire.NewSet(
	userSvc.NewUserService,
	storeSvc.NewStoreService,
	orderSvc.NewOrderService,
	menuSvc.NewMenuService,
	extraSvc.NewExtraService,
)
