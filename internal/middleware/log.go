package middleware

import (
	"os"

	"github.com/labstack/echo/v4/middleware"
)

// LoggerConfig log http request status
var LoggerConfig = middleware.LoggerWithConfig(middleware.LoggerConfig{
	Skipper:          middleware.DefaultSkipper,
	Format:           "${time_custom} ${status} ${method} ${path} (${latency_human})\n",
	CustomTimeFormat: "2006-01-02 15:04:05",
	Output:           os.Stdout,
})
