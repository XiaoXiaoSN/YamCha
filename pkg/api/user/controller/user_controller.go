package controller

import (
	"net/http"
	"yamcha/pkg/api"
	"yamcha/pkg/api/user"

	"github.com/labstack/echo"
)

// UserController is a api controller
type UserController struct {
	userSvc user.Service
}

// NewUsercontroller make a user controllerr
func NewUsercontroller(userSvc user.Service) *UserController {
	return &UserController{
		userSvc: userSvc,
	}
}

// CreateUserEndpoint ...
func (ctl *UserController) CreateUserEndpoint(c echo.Context) error {
	ctx := c.Request().Context()

	u := user.User{}
	err := c.Bind(&u)
	if err != nil {
		return c.JSON(http.StatusCreated, api.H{
			"error": err,
		})
	}

	err = ctl.userSvc.CreateUser(ctx, u)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, api.H{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, api.H{
		"data": true,
	})
}

// UserListEndpoint return users
func (ctl *UserController) UserListEndpoint(c echo.Context) error {
	ctx := c.Request().Context()

	userList, err := ctl.userSvc.UserList(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, api.H{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, api.H{
		"data": userList,
	})
}
