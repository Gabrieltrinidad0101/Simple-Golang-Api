package structs

import (
	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Password string
}

type UserJwt struct {
	Name string `json:name validate:"required"`
	Id   uint   `json:password gorm:"primaryKey" validate:"required"`
	jwt.StandardClaims
}
