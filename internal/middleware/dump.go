package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// BodyDumpConfig log http request body and response body
var BodyDumpConfig = middleware.BodyDumpWithConfig(middleware.BodyDumpConfig{
	Skipper: middleware.DefaultSkipper,
	Handler: func(c echo.Context, reqBody, resBody []byte) {
		// handle request body
		if isJSONContent(c.Request().Header.Get(echo.HeaderContentType)) {
			var prettyJSON bytes.Buffer
			err := json.Indent(&prettyJSON, reqBody, "", "    ")
			if err == nil {
				reqBody = prettyJSON.Bytes()
			}
		}
		// handle response body
		if isJSONContent(c.Response().Header().Get(echo.HeaderContentType)) {
			var prettyJSON bytes.Buffer
			err := json.Indent(&prettyJSON, resBody, "", "    ")
			if err == nil {
				resBody = prettyJSON.Bytes()
			}
		}
		fmt.Printf("request:  %s\nresponse: %s\n\n", reqBody, resBody)
	},
})

func isJSONContent(headerContentType string) bool {
	headerContentType = strings.ToUpper(headerContentType)
	allowList := []string{
		strings.ToUpper(echo.MIMEApplicationJSON),
		strings.ToUpper(echo.MIMEApplicationJSONCharsetUTF8),
	}
	for i := range allowList {
		if headerContentType == allowList[i] {
			return true
		}
	}
	return false
}
