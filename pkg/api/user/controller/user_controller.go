package controller

import (
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
