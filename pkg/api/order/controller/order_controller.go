package controller

import (
	"net/http"
	"strconv"
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

	cParam := order.CreateOrderParams{}
	err := c.Bind(&cParam)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, api.H{"error": err})
	}

	orderObject, err := ctl.orderSvc.CreateOrder(ctx, cParam)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, api.H{"error": err})
	}
	return c.JSON(http.StatusCreated, api.H{
		"data": orderObject,
	})
}

// GetOrderEndpoint return orders
func (ctl *OrderController) GetOrderEndpoint(c echo.Context) error {
	ctx := c.Request().Context()

	orderIDStr := c.Param("orderId")
	orderID, err := strconv.Atoi(orderIDStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, api.H{"error": err})
	}

	orderObject, err := ctl.orderSvc.GetOrder(ctx, orderID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, api.H{"error": err})
	}

	return c.JSON(http.StatusOK, api.H{
		"data": orderObject,
	})
}

// OrderListEndpoint return orders
func (ctl *OrderController) OrderListEndpoint(c echo.Context) error {
	ctx := c.Request().Context()

	findParams := order.Params{}
	err := c.Bind(&findParams)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, api.H{"error": err})
	}

	orderObject, err := ctl.orderSvc.OrderList(ctx, findParams)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, api.H{"error": err})
	}

	return c.JSON(http.StatusOK, api.H{
		"data": orderObject,
	})
}
