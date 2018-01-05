package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

// MainPage test response
func MainPage() echo.HandlerFunc {
	return func(c echo.Context) error {
		param1 := c.Param("param1")
		param2 := c.Param("param2")
		return c.String(http.StatusOK, "test param:"+param1+":"+param2)
	}
}
