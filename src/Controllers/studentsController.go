package controllers

import (
	model "main/src/Model"
	structs "main/src/Structs"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func CreateStudent(ctx echo.Context) error {
	var student structs.Student
	ctx.Bind(student)
	modelStudent := model.NewStudent()
	ok := modelStudent.Create(&student)
	if !ok {
		return ctx.JSON(http.StatusConflict, map[string]string{
			"error": "Error in the DB",
		})
	}
	return ctx.JSON(http.StatusConflict, map[string]structs.Student{
		"data": student,
	})
}

func GetStudents(ctx echo.Context) error {
	modelStudent := model.NewStudent()
	users, ok := modelStudent.Get()
	if !ok {
		return ctx.JSON(http.StatusConflict, map[string]string{
			"error": "Error in the DB",
		})
	}
	return ctx.JSON(http.StatusConflict, map[string][]structs.Student{
		"data": users,
	})
}

func UpdateStudents(ctx echo.Context) error {
	var student structs.Student
	ctx.Bind(student)
	modelStudent := model.NewStudent()
	ok := modelStudent.Update(student)
	if !ok {
		return ctx.JSON(http.StatusConflict, map[string]string{
			"error": "Error in the DB",
		})
	}
	return ctx.JSON(http.StatusConflict, map[string]structs.Student{
		"data": student,
	})
}

func DeleteStudent(ctx echo.Context) error {
	userId, err := strconv.Atoi(ctx.Param("userId"))
	if err != nil {
		return ctx.JSON(http.StatusConflict, map[string]string{
			"error": "Error invalid student id",
		})
	}

	modelStudent := model.NewStudent()
	ok := modelStudent.Delete(uint(userId))
	if !ok {
		return ctx.JSON(http.StatusConflict, map[string]string{
			"error": "Error in the DB",
		})
	}
	return ctx.JSON(http.StatusConflict, map[string]string{
		"message": "User deleted successfully",
	})
}
