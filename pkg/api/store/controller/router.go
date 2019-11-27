package controller

import (
	"github.com/labstack/echo"
)

// SetRoutes regiester stores api
func SetRoutes(engine *echo.Echo, storeCtl *StoreController) {
	apiV1Group := engine.Group("/api/v1")

	apiV1Group.GET("/stores", storeCtl.StoreListEndpoint)
	apiV1Group.GET("/stores/:storeId", storeCtl.BranchStoreListEndpoint)
	apiV1Group.POST("/stores", storeCtl.CreateStoreEndpoint)
	apiV1Group.GET("/stores/:storeId/branchStores", storeCtl.BranchStoreListEndpoint)
	apiV1Group.POST("/stores/:storeId/branchStores", storeCtl.CreateBranchStoreEndpoint)
}
