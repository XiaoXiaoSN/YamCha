package controller

import (
	"github.com/labstack/echo"
	"net/http"
	"strconv"
	"yamcha/pkg/api"
	"yamcha/pkg/api/extra"
)

// ExtraController is a api controller
type ExtraController struct {
	extraSvc extra.Service
}

// NewExtracontroller make a extra controller
func NewExtracontroller(extraSvc extra.Service) *ExtraController {
	return &ExtraController{
		extraSvc: extraSvc,
	}
}

// ExtraListEndpoint get Extra list depends on store ID
func (ctl ExtraController) ExtraListEndpoint(c echo.Context) error {
	ctx := c.Request().Context()

	storeIDStr := c.Param("branchStoreId")
	orderID, err := strconv.Atoi(storeIDStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, api.H{"error": err.Error()})
	}

	orderObject, err := ctl.extraSvc.GetExtraList(ctx, orderID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, api.H{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, api.H{
		"data": orderObject,
	})
}
