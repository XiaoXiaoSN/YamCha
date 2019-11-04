package controller

import (
	"net/http"
	"strconv"
	"yamcha/pkg/api"
	"yamcha/pkg/api/store"

	"github.com/labstack/echo"
)

// StoreController is a api controller
type StoreController struct {
	storeSvc store.Service
}

// NewStorecontroller make a user controllerr
func NewStorecontroller(storeSvc store.Service) *StoreController {
	return &StoreController{
		storeSvc: storeSvc,
	}
}

// CreateStoreEndpoint ...
func (ctl *StoreController) CreateStoreEndpoint(c echo.Context) (err error) {
	ctx := c.Request().Context()

	newStore := store.Store{}
	if err = c.Bind(&newStore); err != nil {
		return c.JSON(http.StatusInternalServerError, api.H{"error": err})
	}

	var storeData store.Store
	storeData, err = ctl.storeSvc.CreateStore(ctx, newStore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, api.H{"error": err})
	}

	return c.JSON(http.StatusCreated, api.H{
		"data": storeData,
	})
}

// StoreListEndpoint return users
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

// BranchStoreListEndpoint return users
func (ctl *StoreController) BranchStoreListEndpoint(c echo.Context) error {
	ctx := c.Request().Context()

	storeIDStr := c.Param("storeId")
	storeID, err := strconv.Atoi(storeIDStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, api.H{"error": err})
	}

	branchStoreList, err := ctl.storeSvc.BranchStoreList(ctx, storeID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, api.H{"error": err})
	}

	return c.JSON(http.StatusOK, api.H{
		"data": branchStoreList,
	})
}
