package controller

import (
	"github.com/labstack/echo"
)

// SetRoutes regiester orders api
func SetRoutes(engine *echo.Echo, menuCtl *MenuController) {
	apiV1Group := engine.Group("/api/v1")

	apiV1Group.GET("/menu/:storeId", menuCtl.MenuListEndpoint)
}
