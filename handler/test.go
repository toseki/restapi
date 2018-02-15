package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

type (
	txJSON struct {
		//DevEUI    string `json:"devEUI"`
		Reference string `json:"ref"`  // reference -> ref FORMAT CHANGE at 2017/10/19
		Confirmed bool   `json:"cnf"`  // confirmed -> cnf
		FPort     uint8  `json:"port"` // fPort -> port
		Data      string `json:"data"`
	}
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

func POSTData(c echo.Context) error {
	txpayload := new(txJSON)
	if err := c.Bind(txpayload); err != nil {
		return err
	}
	//dec := json.NewDecoder(bytes.NewReader(txpayload))
	fmt.Printf("txpayload= %#v\n", txpayload)

	return c.String(http.StatusOK, "txJSON:"+txpayload.Reference)
}
