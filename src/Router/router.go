package router

import (
	controllers "main/src/Controllers"
	middleware "main/src/Middleware"

	"github.com/labstack/echo"
)

func Init(e *echo.Group) {
	//public router
	e.POST("/login", controllers.Login)
	e.POST("/register", controllers.Register)
	//private router
	admin := e.Group("/user")
	admin.Use(middleware.VerifyJwt)
	admin.POST("/create", controllers.CreateStudent)
	admin.GET("/get", controllers.GetStudents)
	admin.PUT("/update/:userId", controllers.UpdateStudents)
	admin.DELETE("/delete/:userId", controllers.DeleteStudent)
}
