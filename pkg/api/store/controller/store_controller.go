package controller

import (
	"net/http"
	"yamcha/pkg/api"
	"yamcha/pkg/api/store"

	"github.com/labstack/echo"
)

// UserController is a api controller
type StoreController struct {
	storeSvc store.Service
}

// NewUsercontroller make a user controllerr
func NewUsercontroller(storeSvc store.Service) *StoreController {
	return &StoreController{
		storeSvc: storeSvc,
	}
}

// CreateUserEndpoint ...
func (ctl *StoreController) CreateStoreEndpoint(c echo.Context) error {
	return nil
}

// UserListEndpoint return users
func (ctl *StoreController) StoreListEndpoint(c echo.Context) error {
	ctx := c.Request().Context()
	storeList, err := ctl.storeSvc.StoreList(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, api.H{"error": err})
	}

	return c.JSON(http.StatusOK, api.H{
		"data": storeList,
	})
}
