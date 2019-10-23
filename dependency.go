package main

import (
	pkgUser "yamcha/pkg/api/user"
	userCtl "yamcha/pkg/api/user/controller"
	userRepo "yamcha/pkg/api/user/repository"
	userSvc "yamcha/pkg/api/user/service"

	pkgStore "yamcha/pkg/api/store"
	storeCtl "yamcha/pkg/api/store/controller"
	storeRepo "yamcha/pkg/api/store/repository"
	storeSvc "yamcha/pkg/api/store/service"

	pkgDB "yamcha/pkg/database"

	_ "github.com/go-sql-driver/mysql"

	"github.com/labstack/echo"
)

var (
	_userRepo pkgUser.Repository
	_userSvc  pkgUser.Service

	_storeRepo pkgStore.Repository
	_storeSvc  pkgStore.Service
)

func initRestfulAPI(e *echo.Echo) error {
	// TODO: config file
	db, err := pkgDB.NewDatabases(pkgDB.Config{
		Username: "xiao",
		Password: "gUKmFVmSdOgTTinmQa9fmYr5AT0EAci5",
		Address:  "yamcha.10oz.tw:23306",
		DBName:   "yamcha_db",
		Env:      "dev",
	})
	if err != nil {
		return nil
	}

	// init Repo
	_userRepo = userRepo.NewUserRepository(db)
	_storeRepo = storeRepo.NewStoreRepository(db)

	// init Service
	_userSvc = userSvc.NewUserService(_userRepo)
	_storeSvc = storeSvc.NewStoreService(_storeRepo)

	// regiest router
	_userCtl := userCtl.NewUsercontroller(_userSvc)
	userCtl.SetRoutes(e, _userCtl)

	_storeCtl := storeCtl.NewUsercontroller(_storeSvc)
	storeCtl.SetRoutes(e, _storeCtl)

	return nil
}
