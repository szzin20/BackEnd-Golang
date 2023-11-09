package middlewares

import (
	"os"
	"healthcare/models/web"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userLoginResponse *web.UserLoginResponse, id int) (string, error) {
	
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["id"] = id
	claims["name"] = userLoginResponse.Fullname
	claims["email"] = userLoginResponse.Email
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	validToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return validToken, nil
}
