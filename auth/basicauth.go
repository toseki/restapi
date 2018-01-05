package auth

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// BasicAuth username password check
func BasicAuth() echo.MiddlewareFunc {
	return middleware.BasicAuth(func(username, password string) bool {
		return username == "test" && password == "test"
	})
}
