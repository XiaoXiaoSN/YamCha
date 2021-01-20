package linebot

import "github.com/labstack/echo/v4"

// SetRoutes register orders api
func SetRoutes(engine *echo.Echo, lineBot LineBot) {
	engine.POST("/callback", lineBot.CallbackHandle)
}
