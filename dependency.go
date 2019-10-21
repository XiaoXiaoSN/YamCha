package main

import (
	pkgUser "yamcha/pkg/api/user"
	userCtl "yamcha/pkg/api/user/controller"
	userRepo "yamcha/pkg/api/user/repository"
	userSvc "yamcha/pkg/api/user/service"

	pkgDB "yamcha/pkg/database"

	_ "github.com/go-sql-driver/mysql"

	"github.com/labstack/echo"
)

var (
	_userRepo pkgUser.Repository
	_userSvc  pkgUser.Service
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

	// init Service
	_userSvc = userSvc.NewUserService(_userRepo)

	// regiest router
	_userCtl := userCtl.NewUsercontroller(_userSvc)
	userCtl.SetRoutes(e, _userCtl)

	return nil
}
