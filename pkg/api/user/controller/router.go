package controller

import (
	"github.com/labstack/echo"
)

// SetRoutes register users api
func SetRoutes(engine *echo.Echo, userCtl *UserController) {
	apiV1Group := engine.Group("/api/v1")

	apiV1Group.GET("/users", userCtl.UserListEndpoint)
	apiV1Group.POST("/users", userCtl.CreateUserEndpoint)
}
