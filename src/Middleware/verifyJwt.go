package middleware

import (
	structs "main/src/Structs"
	"main/src/utils"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func VerifyJwt(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("token")
		if tokenString == "" {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"error": "Invalid token",
			})
		}

		token, err := jwt.ParseWithClaims(tokenString, &structs.UserJwt{}, func(t *jwt.Token) (interface{}, error) {
			return utils.SECRET, nil
		})

		if err != nil {
			return c.String(http.StatusUnauthorized, "Invalid token")
		}

		if _, ok := token.Claims.(*structs.UserJwt); ok && token.Valid {
			c.Set("user", token)
			return next(c)
		}

		return c.String(http.StatusUnauthorized, "Invalid token")
	}
}
