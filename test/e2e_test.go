package test

import (
	"net/http"
	"testing"

	"yamcha/internal/config"
	httpPkg "yamcha/internal/http"

	"github.com/gavv/httpexpect/v2"
)

func testEcho(e *httpexpect.Expect) {
	e.GET("/").
		Expect().
		Status(http.StatusOK).
		Body().Equal("Hello, World!")

	// testUserAPI(e)
}

func TestNewEcho(t *testing.T) {
	cfg := config.NewConfiguration()
	handler := httpPkg.NewEcho(cfg)

	e := httpexpect.WithConfig(httpexpect.Config{
		Client: &http.Client{
			Transport: httpexpect.NewBinder(handler),
			Jar:       httpexpect.NewJar(),
		},
		Reporter: httpexpect.NewAssertReporter(t),
		Printers: []httpexpect.Printer{
			httpexpect.NewDebugPrinter(t, true),
		},
	})

	testEcho(e)
}
