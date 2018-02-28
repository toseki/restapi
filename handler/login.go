package handler

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo"
	"github.com/toseki/restapi/auth"
)

type authJSON struct {
	Token string `json:"token"`
}

func Login(c echo.Context) error {
	username, _, _ := c.Request().BasicAuth() // get BasicAuth username info
	tokenString, err := auth.Gentoken(username)
	if err != nil {
		return c.String(http.StatusNotAcceptable, "Authentication Failure")
	}

	var tokendata authJSON
	tokendata.Token = tokenString
	jsonString, err := json.Marshal(tokendata)
	if err != nil {
		return c.String(http.StatusNotAcceptable, "Authentication Failure")
	}

	return c.String(http.StatusOK, string(jsonString))
}
