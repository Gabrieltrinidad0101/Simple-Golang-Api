package server

import "github.com/labstack/echo"

func StartServer() {
	e := echo.New()

	e.Start(":8080")
}
