package router

import (
	studentsController "main/src/Controllers"

	"github.com/labstack/echo"
)

func Init(e *echo.Group) {
	//public router
	e.POST("/login", studentsController.Login)
	e.POST("/register", studentsController.Register)
	//private router
}
