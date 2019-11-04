package controller

import (
	"github.com/labstack/echo"
)

// SetRoutes regiester orders api
func SetRoutes(engine *echo.Echo, orderCtl *OrderController) {
	apiV1Group := engine.Group("/api/v1")

	apiV1Group.GET("/orders/:orderId", orderCtl.OrderListEndpoint)
	apiV1Group.POST("/orders", orderCtl.CreateOrderEndpoint)
}
