package controller

import (
	"github.com/labstack/echo"
	"net/http"
	"strconv"
	"yamcha/pkg/api"
	"yamcha/pkg/api/menu"
)

// MenuController is a api controller
type MenuController struct {
	menuSvc menu.Service
}

// NewMenucontroller make a menu controller
func NewMenucontroller(menuSvc menu.Service) *MenuController {
	return &MenuController{
		menuSvc: menuSvc,
	}
}

// MenuListEndpoint get Menu list depends on store ID
func (ctl MenuController) MenuListEndpoint(c echo.Context) error {
	ctx := c.Request().Context()

	storeIDStr := c.Param("storeId")
	orderID, err := strconv.Atoi(storeIDStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, api.H{"error": err.Error()})
	}

	orderObject, err := ctl.menuSvc.GetMenuList(ctx, orderID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, api.H{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, api.H{
		"data": orderObject,
	})
}
