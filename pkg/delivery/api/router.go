package api

import "github.com/labstack/echo/v4"

// SetRoutes register orders api
func SetRoutes(engine *echo.Echo, controller Controller) {
	apiV1Group := engine.Group("/api/v1")

	// order
	{
		apiV1Group.POST("/orders", controller.CreateOrderEndpoint)
		// TODO: 這個 path 怎麼會是對應到 list??
		apiV1Group.GET("/orders/:orderId", controller.OrderListEndpoint)
		apiV1Group.PATCH("/orders/:orderId", controller.UpdateOrderEndpoint)
		apiV1Group.DELETE("/orders/:orderId", controller.DeleteOrderEndpoint)
	}

	// store
	{
		apiV1Group.POST("/stores", controller.CreateStoreEndpoint)
		apiV1Group.POST("/stores/:storeId/branchStores", controller.CreateBranchStoreEndpoint)
		apiV1Group.GET("/stores", controller.StoreListEndpoint)
		apiV1Group.GET("/stores/:storeId", controller.BranchStoreListEndpoint)
		apiV1Group.GET("/stores/:storeId/branchStores", controller.BranchStoreListEndpoint)
	}

	// user
	{
		apiV1Group.GET("/users", controller.UserListEndpoint)
		apiV1Group.POST("/users", controller.CreateUserEndpoint)
	}

	// extra
	{
		apiV1Group.GET("/extra/:branchStoreId", controller.ExtraListEndpoint)
	}

	// menu
	{
		apiV1Group.GET("/menu/:branchStoreId", controller.MenuListEndpoint)
	}
}
