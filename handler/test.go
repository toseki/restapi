package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

// MainPage test response
func MainPage() echo.HandlerFunc {
	return func(c echo.Context) error {
		userparam := c.Param("userparam")
		param := c.Param("param")
		username, _, _ := c.Request().BasicAuth() // get BasicAuth username info

		fmt.Printf("userdata: %v,%v\n", userparam, username)

		if username != userparam {
			return c.String(http.StatusNotAcceptable, "not valid user.")
		}

		return c.String(http.StatusOK, "test param:"+userparam+":"+param)
	}
}
