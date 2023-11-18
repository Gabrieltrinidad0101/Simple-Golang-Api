package controllers

import (
	model "main/src/Model"
	structs "main/src/Structs"
	"main/src/utils"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
)

var validate = validator.New()

type StudentsController struct {
}

func calcularCurrentPayment(student structs.Student) (float64, string) {
	if student.BalancePayment <= 0 && student.CurrentPayment >= 0 {
		return 0, "The student cannot have a current payment greater than 0 without a balance"
	}
	if student.BalancePayment > 0 && student.CurrentPayment == 0 {
		return student.BalancePayment * 1.5, ""
	}
	return student.BalancePayment, ""
}

func (s *StudentsController) CreateStudent(ctx echo.Context) error {
	var student structs.Student
	ctx.Bind(&student)
	modelStudent := model.NewStudent()
	if err := validate.Struct(student); err != nil {
		verr, _ := err.(validator.ValidationErrors)
		return ctx.JSON(http.StatusConflict, map[string]string{
			"error": utils.ExtractErrorMessages(verr),
		})
	}

	currentPayment, err := calcularCurrentPayment(student)
	if err != "" {
		return ctx.JSON(http.StatusConflict, structs.ApiResult{
			Error: err,
		})
	}
	student.CurrentPayment = currentPayment
	ok := modelStudent.Create(&student)
	if !ok {
		return ctx.JSON(http.StatusConflict, structs.ApiResult{
			Error: "Error in the DB",
		})
	}
	return ctx.JSON(http.StatusConflict, structs.ApiResult{
		Data: student,
	})
}

func (s *StudentsController) GetStudents(ctx echo.Context) error {
	modelStudent := model.NewStudent()
	users, ok := modelStudent.Get()
	if !ok {
		return ctx.JSON(http.StatusConflict, structs.ApiResult{
			Error: "Error in the DB",
		})
	}

	return ctx.JSON(http.StatusConflict, structs.ApiResult{
		Data: users,
	})
}

func (s *StudentsController) UpdateStudents(ctx echo.Context) error {
	var student structs.Student
	ctx.Bind(&student)
	if err := validate.Struct(student); err != nil {
		verr, _ := err.(validator.ValidationErrors)
		return ctx.JSON(http.StatusConflict, structs.ApiResult{
			Error: utils.ExtractErrorMessages(verr),
		})
	}
	studentId, errorUserId := strconv.Atoi(ctx.Param("studentId"))
	if errorUserId != nil {
		return ctx.JSON(http.StatusConflict, structs.ApiResult{
			Error: "Error invalid student id",
		})
	}
	modelStudent := model.NewStudent()
	currentPayment, err := calcularCurrentPayment(student)
	if err != "" {
		return ctx.JSON(http.StatusConflict, structs.ApiResult{
			Error: err,
		})
	}
	student.CurrentPayment = currentPayment
	ok := modelStudent.Update(uint(studentId), &student)
	if !ok {
		return ctx.JSON(http.StatusConflict, structs.ApiResult{
			Error: "Error in the DB",
		})
	}
	return ctx.JSON(http.StatusConflict, structs.ApiResult{
		Data: student,
	})
}

func (s *StudentsController) DeleteStudent(ctx echo.Context) error {
	userId, err := strconv.Atoi(ctx.Param("userId"))
	if err != nil {
		return ctx.JSON(http.StatusConflict, structs.ApiResult{
			Error: "Error invalid student id",
		})
	}

	modelStudent := model.NewStudent()
	ok := modelStudent.Delete(uint(userId))
	if !ok {
		return ctx.JSON(http.StatusConflict, structs.ApiResult{
			Error: "Error in the DB",
		})
	}
	return ctx.JSON(http.StatusConflict, structs.ApiResult{
		Message: "User deleted successfully",
	})
}
