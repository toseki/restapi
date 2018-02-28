package handler

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/toseki/restapi/auth"
)

type (
	txJSON struct {
		Field1 string `json:"field1"`
		Field2 bool   `json:"field2"`
		Field3 uint8  `json:"field3"`
		Field4 string `json:"field4"`
	}

	apiauth struct {
		Jwt string `json:"jwt"`
	}

	node struct {
		AppFiled string `json:"appField"`
	}
)

// MainPage test response
func MainPage() echo.HandlerFunc {
	return func(c echo.Context) error {
		userparam := c.Param("userparam")
		param := c.Param("param")
		//username, _, _ := c.Request().BasicAuth() // get BasicAuth username info
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(*auth.UserClaims)
		username := claims.Username
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

	// TEST POST get AuthToken start
	jsonStr := `{"password":"` + `password` + `","username":"` + `user` + `"}` // sample:username:password
	url := "https://localhost:8088/api/hoge/login"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(jsonStr)))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	//defer resp.Body.Close()

	fmt.Printf("resp=%s\n", resp.Body)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	var apitoken apiauth
	err = json.Unmarshal(body, &apitoken)
	if err != nil {
		return err
	}
	fmt.Printf("Jwt=%s\n", apitoken.Jwt)
	resp.Body.Close()
	//

	// TEST GET get DevEUI's appid
	url = "https://localhost:8088/api/hogenode/" + "hogedevid"
	req, err = http.NewRequest("GET", url, bytes.NewBuffer([]byte(jsonStr)))
	if err != nil {
		return err
	}

	req.Header.Set("AuthToken", apitoken.Jwt)

	resp, err = client.Do(req)
	if err != nil {
		return err
	}

	fmt.Printf("resp=%s\n", resp.Body)
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var nodeinfo node
	err = json.Unmarshal(body, &nodeinfo)
	if err != nil {
		return err
	}
	fmt.Printf("AppField=%s\n", nodeinfo.AppFiled)
	resp.Body.Close()

	return c.String(http.StatusOK, "txJSON:"+txpayload.Field1+"resp="+fmt.Sprintf("%d", resp.StatusCode)+":"+apitoken.Jwt+" AppField="+nodeinfo.AppFiled)
}
