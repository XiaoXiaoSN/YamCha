package controller

import (
	"github.com/labstack/echo"
)

// SetRoutes regiester users api
func SetRoutes(engine *echo.Echo, userCtl *UserController) {
	engine.POST("/users", userCtl.CreateUserEndpoint)
}
