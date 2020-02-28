package main

import (
	"net/http"
	"os"

	pkgUser "yamcha/pkg/api/user"
	userCtl "yamcha/pkg/api/user/controller"
	userRepo "yamcha/pkg/api/user/repository"
	userSvc "yamcha/pkg/api/user/service"

	pkgStore "yamcha/pkg/api/store"
	storeCtl "yamcha/pkg/api/store/controller"
	storeRepo "yamcha/pkg/api/store/repository"
	storeSvc "yamcha/pkg/api/store/service"

	pkgOrder "yamcha/pkg/api/order"
	orderCtl "yamcha/pkg/api/order/controller"
	orderRepo "yamcha/pkg/api/order/repository"
	orderSvc "yamcha/pkg/api/order/service"

	pkgMenu "yamcha/pkg/api/menu"
	menuCtl "yamcha/pkg/api/menu/controller"
	menuRepo "yamcha/pkg/api/menu/repository"
	menuSvc "yamcha/pkg/api/menu/service"

	pkgExtra "yamcha/pkg/api/extra"
	extraCtl "yamcha/pkg/api/extra/controller"
	extraRepo "yamcha/pkg/api/extra/repository"
	extraSvc "yamcha/pkg/api/extra/service"

	"yamcha/pkg/linebot"

	pkgConfig "yamcha/internal/pkg/config"
	pkgDB "yamcha/internal/pkg/database"

	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var (
	lineChannelSecret string
	lineChannelToken  string

	bot linebot.LineBot

	_userRepo pkgUser.Repository
	_userSvc  pkgUser.Service

	_storeRepo pkgStore.Repository
	_storeSvc  pkgStore.Service

	_orderRepo pkgOrder.Repository
	_orderSvc  pkgOrder.Service

	_menuRepo pkgMenu.Repository
	_menuSvc  pkgMenu.Service

	_extraRepo pkgExtra.Repository
	_extraSvc  pkgExtra.Service
)

var middlewareCfg = middleware.CORSConfig{
	AllowOrigins: []string{"*"},
	AllowMethods: []string{
		http.MethodGet,
		http.MethodPost,
		http.MethodPut,
		http.MethodDelete,
		http.MethodPatch,
	},
	AllowHeaders: []string{
		"*",
		echo.HeaderAuthorization,
		echo.HeaderContentType,
		echo.HeaderOrigin,
		echo.HeaderContentLength,
	},
}

func init() {
	lineChannelSecret = os.Getenv("LINECORP_PLATFORM_CHANNEL_CHANNELSECRET")
	lineChannelToken = os.Getenv("LINECORP_PLATFORM_CHANNEL_CHANNELTOKEN")
}

func initService(e *echo.Echo, cfg *pkgConfig.Configuration) (err error) {
	log.Info("start to init service...")

	// init dependency services
	err = initDependencyService(e, cfg)
	if err != nil {
		return err
	}

	// init yamcha bot
	bot, err = linebot.NewYambotLineBot(lineChannelSecret, lineChannelToken, _orderSvc)
	if err != nil {
		log.Infof("failed to init linebot.NewYambotLineBot err: %+v", err)
		return err
	}

	// register restful API
	e.Use(middleware.CORSWithConfig(middlewareCfg))
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.POST("/callback", bot.CallbackHandle)

	return nil
}

func initDependencyService(e *echo.Echo, cfg *pkgConfig.Configuration) error {
	db, err := pkgDB.NewDatabases(cfg.DBCfg)
	if err != nil {
		return nil
	}

	// init Repo
	_userRepo = userRepo.NewUserRepository(db)
	_storeRepo = storeRepo.NewStoreRepository(db)
	_orderRepo = orderRepo.NewOrderRepository(db)
	_menuRepo = menuRepo.NewMenuRepository(db)
	_extraRepo = extraRepo.NewExtraRepository(db)

	// init Service
	_userSvc = userSvc.NewUserService(_userRepo)
	_storeSvc = storeSvc.NewStoreService(_storeRepo)
	_orderSvc = orderSvc.NewOrderService(_orderRepo)
	_menuSvc = menuSvc.NewMenuService(_menuRepo)
	_extraSvc = extraSvc.NewExtraService(_extraRepo)

	// regiest router
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
