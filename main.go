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
	//e.Use(auth.BasicAuth())
	login := e.Group("/login")
	login.Use(auth.BasicAuth())
	login.GET("", handler.Login)

	//e.GET("/login", handler.Login)

	jwttest := e.Group("/test")

	jwttest.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:     &auth.UserClaims{},
		SigningKey: []byte(auth.Secret),
	}))
	jwttest.GET("/:userparam/:param", handler.MainPage())
	jwttest.GET("/tx", handler.POSTData)

	//e.GET("/test/:userparam/:param", handler.MainPage())
	//e.POST("/test/tx", handler.POSTData)

	e.Start(":8080")
}
