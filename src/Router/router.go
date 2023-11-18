package router

import (
	controllers "main/src/Controllers"
	middleware "main/src/Middleware"

	"github.com/labstack/echo"
)

func Init(e *echo.Group) {
	//public router
	userController := controllers.UserController{}
	e.POST("/login", userController.Login)
	e.POST("/register", userController.Register)

	//private router
	studentController := controllers.StudentsController{}
	admin := e.Group("/user")
	admin.Use(middleware.VerifyJwt)
	admin.POST("/create", studentController.CreateStudent)
	admin.GET("/get", studentController.GetStudents)
	admin.PUT("/update/:userId", studentController.UpdateStudents)
	admin.DELETE("/delete/:userId", studentController.DeleteStudent)
}
