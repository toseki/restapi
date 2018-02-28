package auth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

const Secret = "CWtf85m7fahh0qYTxYaDV3eMpRMjhAO37o3wz/y4IMQ=" // openssl rand -base64 32

type UserClaims struct {
	Username string
	jwt.StandardClaims
}

func Gentoken(username string) (string, error) {

	claims := &UserClaims{
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 3).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(Secret))
	if err != nil {
		return tokenString, err
	}
	return tokenString, nil
}

/* not use
func chktoken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	fmt.Printf("chktoken token=%v\n", token)

	var user UserClaims
	token, err = jwt.ParseWithClaims(tokenString, &user, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	fmt.Printf("chktoken token parse struct=%v\n", user)
	if err != nil {
		return user.Username, err
	}
	return user.Username, nil
}
*/
