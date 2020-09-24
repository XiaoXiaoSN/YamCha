// 不加上去 wire 會錯，我不知道為什麼...
//+build !wireinject

package test

import (
	"net/http"
	"testing"

	"yamcha/cmd/line"
	"yamcha/internal/config"
	httpPkg "yamcha/internal/http"

	"github.com/gavv/httpexpect/v2"
	_ "github.com/go-sql-driver/mysql"
)

func testEcho(e *httpexpect.Expect) {
	e.GET("/").
		Expect().
		Status(http.StatusOK).
		Body().Equal("Hello, World!")

	testUserAPI(e)
}

func TestNewEcho(t *testing.T) {
	cfg := config.NewConfiguration()
	handler := httpPkg.NewEcho(cfg)
	err := line.InitDependencyService(handler, cfg)
	if err != nil {
		t.Error(err)
		return
	}

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
