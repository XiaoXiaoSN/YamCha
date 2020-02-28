package controller

import (
	"github.com/labstack/echo"
)

// SetRoutes register orders api
func SetRoutes(engine *echo.Echo, extraCtl *ExtraController) {
	apiV1Group := engine.Group("/api/v1")

	apiV1Group.GET("/extra/:branchStoreId", extraCtl.ExtraListEndpoint)
}
