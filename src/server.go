package server

import (
	router "main/src/Router"

	"github.com/labstack/echo"
)

func StartServer() {
	e := echo.New()
	router.Init(e.Group(""))
	e.Start(":8080")
}
