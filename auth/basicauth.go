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
		/* test jwt
		tokenString := c.Request().Header.Get("token")
		if tokenString != "" {
			tokenUsername, err := chktoken(tokenString)
			if err != nil {
				fmt.Printf("BasicAuth: chktoken err:%s", err)
				return false, nil
			}
			fmt.Printf("BasicAuth: chktoken username=%s", tokenUsername)
			return true, nil
		}
		*/

		salt := "salt" // sample salt string
		//passwd := "password"
		//saltpasswd := sha256.Sum256([]byte(salt + passwd))
		sha256hex := "13601bda4ea78e55a07b98866d2be6be0744e3866f13c00c811cab608a28f322" // sha256 hash hexdigest salt+password
		authchkhex := fmt.Sprintf("%x", sha256.Sum256([]byte(salt+password)))

		fmt.Println("saltpassword:", authchkhex)

		return username == "user" && sha256hex == authchkhex, nil
	})
}
