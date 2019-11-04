package controller

import (
	"net/http"
	"yamcha/pkg/api"
	"yamcha/pkg/api/order"

	"github.com/labstack/echo"
)

// OrderController is a api controller
type OrderController struct {
	orderSvc order.Service
}

// NewOrdercontroller make a order controllerr
func NewOrdercontroller(orderSvc order.Service) *OrderController {
	return &OrderController{
		orderSvc: orderSvc,
	}
}

// CreateOrderEndpoint ...
func (ctl *OrderController) CreateOrderEndpoint(c echo.Context) error {
	ctx := c.Request().Context()

	param := &order.Params{}
	e := c.Bind(param)
	if e == nil {
		orderObject, err := ctl.orderSvc.CreateOrder(ctx, *param)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, api.H{"error": err})
		}
		return c.JSON(http.StatusOK, api.H{
			"data": orderObject,
		})
	}
	return c.JSON(http.StatusInternalServerError, api.H{"error": e})
	// id := c.Param("channelId")

}

// OrderListEndpoint return orders
func (ctl *OrderController) OrderListEndpoint(c echo.Context) error {
	ctx := c.Request().Context()
	id := c.Param("orderId")
	orderObject, err := ctl.orderSvc.OrderList(ctx, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, api.H{"error": err})
	}

	return c.JSON(http.StatusOK, api.H{
		"data": orderObject,
	})
}
