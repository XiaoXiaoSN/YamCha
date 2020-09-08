package controller

import (
	"log"
	"net/http"
	"strconv"
	"yamcha/pkg/api"
	"yamcha/pkg/api/store"

	"github.com/labstack/echo/v4"
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
		return c.JSON(http.StatusInternalServerError, api.H{"error": err.Error()})
	}

	var storeData store.Store
	storeData, err = ctl.storeSvc.CreateStore(ctx, newStore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, api.H{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, api.H{
		"data": storeData,
	})
}

// StoreListEndpoint all of stores
func (ctl *StoreController) StoreListEndpoint(c echo.Context) error {
	ctx := c.Request().Context()

	storeList, err := ctl.storeSvc.StoreList(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, api.H{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, api.H{
		"data": storeList,
	})
}

// GetStoreEndpoint return target store
func (ctl *StoreController) GetStoreEndpoint(c echo.Context) error {
	ctx := c.Request().Context()

	storeIDStr := c.Param("storeId")
	storeID, err := strconv.Atoi(storeIDStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, api.H{"error": err.Error()})
	}

	targetStore, err := ctl.storeSvc.GetStore(ctx, storeID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, api.H{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, api.H{
		"data": targetStore,
	})
}

// BranchStoreListEndpoint return users
func (ctl *StoreController) BranchStoreListEndpoint(c echo.Context) error {
	ctx := c.Request().Context()

	storeIDStr := c.Param("storeId")
	storeID, err := strconv.Atoi(storeIDStr)
	log.Println("in Branch:", storeID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, api.H{"error": err.Error()})
	}

	branchStoreList, err := ctl.storeSvc.BranchStoreList(ctx, storeID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, api.H{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, api.H{
		"data": branchStoreList,
	})
}

// CreateBranchStoreEndpoint ...
func (ctl *StoreController) CreateBranchStoreEndpoint(c echo.Context) (err error) {
	ctx := c.Request().Context()

	newBranchStore := store.BranchStore{}
	if err = c.Bind(&newBranchStore); err != nil {
		return c.JSON(http.StatusInternalServerError, api.H{"error": err.Error()})
	}

	storeIDStr := c.Param("storeId")
	storeID, err := strconv.Atoi(storeIDStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, api.H{"error": err.Error()})
	}
	newBranchStore.StoreGroupID = storeID

	var storeData store.BranchStore
	storeData, err = ctl.storeSvc.CreateBranchStore(ctx, newBranchStore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, api.H{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, api.H{
		"data": storeData,
	})
}
