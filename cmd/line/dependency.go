package line

import (
	pkgUser "yamcha/pkg/service/user"
	userSvc "yamcha/pkg/service/user/service"

	pkgStore "yamcha/pkg/service/store"
	storeSvc "yamcha/pkg/service/store/service"

	pkgOrder "yamcha/pkg/service/order"
	orderSvc "yamcha/pkg/service/order/service"

	pkgMenu "yamcha/pkg/service/menu"
	menuSvc "yamcha/pkg/service/menu/service"

	pkgExtra "yamcha/pkg/service/extra"
	extraSvc "yamcha/pkg/service/extra/service"

	"yamcha/pkg/delivery/api"

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
	ctrl := api.NewController(_orderSvc, _storeSvc, _userSvc, _menuSvc, _extraSvc)
	api.SetRoutes(e, ctrl)

	return nil
}
