package auth

import (
	"crypto/sha256"
	"fmt"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// BasicAuth username password check
func BasicAuth() echo.MiddlewareFunc {
	return middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		salt := "salt" // sample salt string
		//passwd := "password"
		//saltpasswd := sha256.Sum256([]byte(salt + passwd))
		sha256hex := "13601bda4ea78e55a07b98866d2be6be0744e3866f13c00c811cab608a28f322"
		authchkhex := fmt.Sprintf("%x", sha256.Sum256([]byte(salt+password)))

		fmt.Println("saltpassword:", authchkhex)

		return username == "test-1" && sha256hex == authchkhex, nil
	})
}
