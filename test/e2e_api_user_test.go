// 不加上去 wire 會錯，我不知道為什麼...
//+build !wireinject

package test

import (
	"net/http"

	"yamcha/pkg/api/user"

	"github.com/gavv/httpexpect/v2"
)

func testUserAPI(e *httpexpect.Expect) {
	e.POST("/api/v1/users").
		Expect().
		Status(http.StatusBadRequest)

	u := user.User{
		Name:   "Testing",
		LineID: "55688",
	}
	e.POST("/api/v1/users").WithJSON(&u).
		Expect().
		Status(http.StatusCreated)
}
