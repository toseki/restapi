package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/toseki/restapi/auth"
	"github.com/toseki/restapi/handler"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(auth.BasicAuth())

	e.GET("/test/:param1/:param2", handler.MainPage())

	e.Start(":8080")
}
