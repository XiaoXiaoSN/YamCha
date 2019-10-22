package controller

import (
	"net/http"
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
	return nil
}

// UserListEndpoint return users
func (ctl *UserController) UserListEndpoint(c echo.Context) error {
	ctx := c.Request().Context()

	userList, err := ctl.userSvc.UserList(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, userList)
}
