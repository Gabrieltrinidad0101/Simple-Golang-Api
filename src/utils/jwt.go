package utils

import (
	structs "main/src/Structs"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const SECRET = "prueba"

func CreateJsonWebToken(user structs.User) (tokenString string, err error) {
	userJwt := &structs.UserJwt{
		Name: user.Name,
		Id:   user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userJwt)
	tokenString, err = token.SignedString(SECRET)

	return
}
