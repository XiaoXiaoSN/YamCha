package line

import (
	pkgUser "yamcha/pkg/api/user"
	userCtl "yamcha/pkg/api/user/controller"
	userSvc "yamcha/pkg/api/user/service"

	pkgStore "yamcha/pkg/api/store"
	storeCtl "yamcha/pkg/api/store/controller"
	storeSvc "yamcha/pkg/api/store/service"

	pkgOrder "yamcha/pkg/api/order"
	orderCtl "yamcha/pkg/api/order/controller"
	orderSvc "yamcha/pkg/api/order/service"

	pkgMenu "yamcha/pkg/api/menu"
	menuCtl "yamcha/pkg/api/menu/controller"
	menuSvc "yamcha/pkg/api/menu/service"

	pkgExtra "yamcha/pkg/api/extra"
	extraCtl "yamcha/pkg/api/extra/controller"
	extraSvc "yamcha/pkg/api/extra/service"

	"yamcha/pkg/linebot"
	"yamcha/pkg/repository"
	dbRepo "yamcha/pkg/repository/db"

	pkgConfig "yamcha/internal/config"
	pkgDB "yamcha/internal/database"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

var (
	bot linebot.LineBot

	_userSvc  pkgUser.Service
	_storeSvc pkgStore.Service
	_orderSvc pkgOrder.Service
	_menuSvc  pkgMenu.Service
	_extraSvc pkgExtra.Service

	repo repository.Repository
)

func initService(e *echo.Echo, cfg *pkgConfig.Configuration) (err error) {
	log.Info("start to init service...")

	// init dependency services
	err = InitDependencyService(e, cfg)
	if err != nil {
		return err
	}

	// init yamcha bot
	bot, err = linebot.NewYambotLineBot(cfg.BotCfg, _orderSvc)
	if err != nil {
		log.Infof("failed to init linebot.NewYambotLineBot err: %+v", err)
		return err
	}

	// register restful API
	e.POST("/callback", bot.CallbackHandle)

	return nil
}

// InitDependencyService Init dependency services
func InitDependencyService(e *echo.Echo, cfg *pkgConfig.Configuration) error {
	db, err := pkgDB.NewDatabases(cfg.DBCfg)
	if err != nil {
		return nil
	}

	// init Repo
	repo = dbRepo.NewRepo(db)

	// init Service
	_userSvc = userSvc.NewUserService(repo)
	_storeSvc = storeSvc.NewStoreService(repo)
	_orderSvc = orderSvc.NewOrderService(repo)
	_menuSvc = menuSvc.NewMenuService(repo)
	_extraSvc = extraSvc.NewExtraService(repo)

	// register router
	_userCtl := userCtl.NewUsercontroller(_userSvc)
	userCtl.SetRoutes(e, _userCtl)

	_storeCtl := storeCtl.NewStorecontroller(_storeSvc)
	storeCtl.SetRoutes(e, _storeCtl)

	_orderCtl := orderCtl.NewOrdercontroller(_orderSvc)
	orderCtl.SetRoutes(e, _orderCtl)

	_menuCtl := menuCtl.NewMenucontroller(_menuSvc)
	menuCtl.SetRoutes(e, _menuCtl)

	_extraCtl := extraCtl.NewExtracontroller(_extraSvc)
	extraCtl.SetRoutes(e, _extraCtl)

	return nil
}
