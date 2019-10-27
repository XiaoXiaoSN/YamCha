package controller

import (
	"net/http"
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
func (ctl *StoreController) CreateStoreEndpoint(c echo.Context) error {
	ctx := c.Request().Context()
	data := &store.Store{}
	e := c.Bind(data)
	if e == nil {
		storeData, err := ctl.storeSvc.CreateStore(ctx, *data)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, api.H{"error": err})
		} else {
			return c.JSON(http.StatusOK, api.H{
				"data": storeData,
			})
		}
	} else {
		return c.JSON(http.StatusInternalServerError, api.H{"error": e})
	}

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
