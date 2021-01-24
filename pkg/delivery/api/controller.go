package api

import (
	"log"
	"net/http"
	"strconv"

	"yamcha/internal/httputil"
	"yamcha/pkg/model"
	"yamcha/pkg/service/extra"
	"yamcha/pkg/service/menu"
	"yamcha/pkg/service/order"
	"yamcha/pkg/service/store"
	"yamcha/pkg/service/user"

	"github.com/labstack/echo/v4"
)

// Controller ...
type Controller interface {
	CreateOrderEndpoint(c echo.Context) error
	GetOrderEndpoint(c echo.Context) error
	OrderListEndpoint(c echo.Context) error
	UpdateOrderEndpoint(c echo.Context) error
	DeleteOrderEndpoint(c echo.Context) error
	CreateStoreEndpoint(c echo.Context) (err error)

	CreateBranchStoreEndpoint(c echo.Context) (err error)
	StoreListEndpoint(c echo.Context) error
	GetStoreEndpoint(c echo.Context) error
	BranchStoreListEndpoint(c echo.Context) error

	CreateUserEndpoint(c echo.Context) error
	UserListEndpoint(c echo.Context) error

	ExtraListEndpoint(c echo.Context) error

	MenuListEndpoint(c echo.Context) error
}

type _controller struct {
	orderSvc order.Service
	storeSvc store.Service
	userSvc  user.Service
	menuSvc  menu.Service
	extraSvc extra.Service
}

// NewController create a api controller
func NewController(
	orderSvc order.Service,
	storeSvc store.Service,
	userSvc user.Service,
	menuSvc menu.Service,
	extraSvc extra.Service,
) Controller {
	return &_controller{
		orderSvc: orderSvc,
		storeSvc: storeSvc,
		userSvc:  userSvc,
		menuSvc:  menuSvc,
		extraSvc: extraSvc,
	}
}

// CreateOrderEndpoint ...
// POST /api/v1/orders
func (ctl *_controller) CreateOrderEndpoint(c echo.Context) error {
	ctx := c.Request().Context()

	cParam := model.CreateOrderParams{}
	err := c.Bind(&cParam)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, httputil.H{"error": err.Error()})
	}
	orderObject, err := ctl.orderSvc.CreateOrder(ctx, cParam)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, httputil.H{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, httputil.H{
		"data": orderObject,
	})
}

// GetOrderEndpoint return orders
func (ctl *_controller) GetOrderEndpoint(c echo.Context) error {
	ctx := c.Request().Context()

	orderIDStr := c.Param("orderId")
	orderID, err := strconv.Atoi(orderIDStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, httputil.H{"error": err.Error()})
	}

	orderObject, err := ctl.orderSvc.GetOrder(ctx, orderID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, httputil.H{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, httputil.H{
		"data": orderObject,
	})
}

// OrderListEndpoint return orders
// GET /api/v1/orders/:orderId
func (ctl *_controller) OrderListEndpoint(c echo.Context) error {
	ctx := c.Request().Context()

	findParams := model.OrderParams{}
	err := c.Bind(&findParams)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, httputil.H{"error": err.Error()})
	}

	orderObject, err := ctl.orderSvc.OrderList(ctx, findParams)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, httputil.H{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, httputil.H{
		"data": orderObject,
	})
}

// UpdateOrderEndpoint ...
// PATCH /api/v1/orders/:orderId
func (ctl *_controller) UpdateOrderEndpoint(c echo.Context) error {
	ctx := c.Request().Context()

	cParam := model.CreateOrderParams{}
	err := c.Bind(&cParam)
	log.Println(cParam)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, httputil.H{"error": err.Error()})
	}
	orderObject, err := ctl.orderSvc.UpdateOrder(ctx, cParam)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, httputil.H{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, httputil.H{
		"data": orderObject,
	})
}

// DeleteOrderEndpoint handle order delete logic
// DELETE /api/v1orders/:orderId
func (ctl *_controller) DeleteOrderEndpoint(c echo.Context) error {
	ctx := c.Request().Context()

	orderIDStr := c.Param("orderId")
	orderID, err := strconv.Atoi(orderIDStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, httputil.H{"error": err.Error()})
	}

	if errorMsg := ctl.orderSvc.DeleteOrder(ctx, orderID); errorMsg != nil {
		return c.JSON(http.StatusInternalServerError, httputil.H{"error": errorMsg.Error()})
	}

	return c.JSON(http.StatusOK, httputil.H{
		"ok": "200",
	})
}

// CreateStoreEndpoint ...
// POST /api/v1/stores
func (ctl *_controller) CreateStoreEndpoint(c echo.Context) (err error) {
	ctx := c.Request().Context()

	newStore := model.Store{}
	if err = c.Bind(&newStore); err != nil {
		return c.JSON(http.StatusInternalServerError, httputil.H{"error": err.Error()})
	}

	var storeData model.Store
	storeData, err = ctl.storeSvc.CreateStore(ctx, newStore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, httputil.H{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, httputil.H{
		"data": storeData,
	})
}

// CreateBranchStoreEndpoint ...
// POST /api/v1/stores/:storeId/branchStores
func (ctl *_controller) CreateBranchStoreEndpoint(c echo.Context) (err error) {
	ctx := c.Request().Context()

	newBranchStore := model.BranchStore{}
	if err = c.Bind(&newBranchStore); err != nil {
		return c.JSON(http.StatusInternalServerError, httputil.H{"error": err.Error()})
	}

	storeIDStr := c.Param("storeId")
	_, err = strconv.Atoi(storeIDStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, httputil.H{"error": err.Error()})
	}
	newBranchStore.StoreGroupID = storeIDStr

	var storeData model.BranchStore
	storeData, err = ctl.storeSvc.CreateBranchStore(ctx, newBranchStore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, httputil.H{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, httputil.H{
		"data": storeData,
	})
}

// StoreListEndpoint list all of stores
// GET /api/v1/stores
func (ctl *_controller) StoreListEndpoint(c echo.Context) error {
	ctx := c.Request().Context()

	storeList, err := ctl.storeSvc.StoreList(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, httputil.H{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, httputil.H{
		"data": storeList,
	})
}

// GetStoreEndpoint return target store
// GET /api/v1/stores/:storeId
func (ctl *_controller) GetStoreEndpoint(c echo.Context) error {
	ctx := c.Request().Context()

	storeIDStr := c.Param("storeId")
	storeID, err := strconv.Atoi(storeIDStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, httputil.H{"error": err.Error()})
	}

	targetStore, err := ctl.storeSvc.GetStore(ctx, storeID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, httputil.H{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, httputil.H{
		"data": targetStore,
	})
}

// BranchStoreListEndpoint return users
// GET /api/v1/stores/:storeId
// GET /api/v1/stores/:storeId/branchStores
func (ctl *_controller) BranchStoreListEndpoint(c echo.Context) error {
	ctx := c.Request().Context()

	storeIDStr := c.Param("storeId")
	storeID, err := strconv.Atoi(storeIDStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, httputil.H{"error": err.Error()})
	}

	branchStoreList, err := ctl.storeSvc.BranchStoreList(ctx, storeID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, httputil.H{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, httputil.H{
		"data": branchStoreList,
	})
}

// CreateUserEndpoint ...
// POST /api/v1/users
func (ctl *_controller) CreateUserEndpoint(c echo.Context) error {
	ctx := c.Request().Context()

	u := model.User{}
	err := c.Bind(&u)
	if err != nil {
		return c.JSON(http.StatusBadRequest, httputil.H{"error": err.Error()})
	}
	err = c.Validate(u)
	if err != nil {
		return c.JSON(http.StatusBadRequest, httputil.H{"error": err.Error()})
	}

	err = ctl.userSvc.CreateUser(ctx, u)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, httputil.H{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, httputil.H{
		"data": true,
	})
}

// UserListEndpoint return users
// GET /api/v1/users
func (ctl *_controller) UserListEndpoint(c echo.Context) error {
	ctx := c.Request().Context()

	userList, err := ctl.userSvc.UserList(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, httputil.H{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, httputil.H{
		"data": userList,
	})
}

// ExtraListEndpoint get Extra list depends on store ID
// GET /api/v1/extra/:branchStoreId
func (ctl *_controller) ExtraListEndpoint(c echo.Context) error {
	ctx := c.Request().Context()

	storeIDStr := c.Param("branchStoreId")
	orderID, err := strconv.Atoi(storeIDStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, httputil.H{"error": err.Error()})
	}

	orderObject, err := ctl.extraSvc.GetExtraList(ctx, orderID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, httputil.H{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, httputil.H{
		"data": orderObject,
	})
}

// MenuListEndpoint get Menu list depends on store ID
// GET /api/v1/menu/:branchStoreId
func (ctl *_controller) MenuListEndpoint(c echo.Context) error {
	ctx := c.Request().Context()

	storeIDStr := c.Param("branchStoreId")
	log.Println(storeIDStr)
	orderID, err := strconv.Atoi(storeIDStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, httputil.H{"error": err.Error()})
	}

	orderObject, err := ctl.menuSvc.GetMenuList(ctx, orderID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, httputil.H{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, httputil.H{
		"data": orderObject,
	})
}
