package controller

import (
	"github.com/labstack/echo"
)

// SetRoutes regiester orders api
func SetRoutes(engine *echo.Echo, orderCtl *OrderController) {
	engine.GET("/orders/:orderId", orderCtl.OrderListEndpoint)
	engine.POST("/orders", orderCtl.CreateOrderEndpoint)
}
