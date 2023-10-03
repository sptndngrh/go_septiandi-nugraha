package helper

import (
	"cleanandhexa_archi/constant"
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateToken(email string) (string, error) {
	claims := jwt.MapClaims{}
	claims["email"] = email
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(10 * time.Minute).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(constant.SECRET_JWT))
}
