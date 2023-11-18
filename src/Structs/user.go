package structs

import (
	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `validate:"required"`
	Password string `validate:"required"`
}

type UserJwt struct {
	Name string `json:"name"`
	Id   uint   `json:"password"`
	jwt.StandardClaims
}
