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
	return nil
}

// OrderListEndpoint return orders
func (ctl *OrderController) OrderListEndpoint(c echo.Context) error {
	ctx := c.Request().Context()

	orderList, err := ctl.orderSvc.OrderList(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, api.H{"error": err})
	}

	return c.JSON(http.StatusOK, api.H{
		"data": orderList,
	})
}
