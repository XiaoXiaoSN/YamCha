package http

import (
	"net/http"
	"yamcha/internal/config"
	"yamcha/internal/middleware"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

// CustomValidator ...
type CustomValidator struct {
	validator *validator.Validate
}

// Validate ...
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

// NewEcho create new engine for handler to register
func NewEcho(cfg *config.Configuration) *echo.Echo {
	e := echo.New()

	// register echo middleware
	e.Use(middleware.CORSConfig)

	if cfg.Env != "production" {
		// log http request body and response body
		e.Use(middleware.BodyDumpConfig)

		// log http request status
		e.Use(middleware.LoggerConfig)
	}

	// register validator
	e.Validator = &CustomValidator{validator: validator.New()}

	RegisterDefaultRoute(e)
	return e
}

// RegisterDefaultRoute provide default handler
func RegisterDefaultRoute(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong!!!")
	})
}
