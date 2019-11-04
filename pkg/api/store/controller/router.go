package controller

import (
	"github.com/labstack/echo"
)

// SetRoutes regiester stores api
func SetRoutes(engine *echo.Echo, storeCtl *StoreController) {
	engine.GET("/stores", storeCtl.StoreListEndpoint)
	engine.GET("/stores/:storeId", storeCtl.BranchStoreListEndpoint)
	engine.POST("/stores", storeCtl.CreateStoreEndpoint)
}
