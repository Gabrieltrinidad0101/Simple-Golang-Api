package controllers

import (
	model "main/src/Model"
	structs "main/src/Structs"
	"main/src/utils"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
)

type UserController struct {
}

func (u *UserController) Login(ctx echo.Context) error {
	var user structs.User
	ctx.Bind(&user)
	if err := validate.Struct(user); err != nil {
		verr, _ := err.(validator.ValidationErrors)
		return ctx.JSON(http.StatusConflict, structs.ApiResult{
			Error: utils.ExtractErrorMessages(verr),
		})
	}
	modelUser := model.NewUser()
	existUser, ok := modelUser.Find(user.Name, user.Password)

	if !ok {
		return ctx.JSON(http.StatusConflict, structs.ApiResult{
			Error: "Error in the DB",
		})
	}

	if existUser.Name == "" {
		return ctx.JSON(http.StatusConflict, structs.ApiResult{
			Error: "The name or password is incorrect",
		})
	}

	token, err := utils.CreateJsonWebToken(existUser)

	if err != nil {
		return ctx.JSON(http.StatusConflict, structs.ApiResult{
			Error: "Error creating the jwt",
		})
	}

	return ctx.JSON(http.StatusOK, structs.ApiResult{
		Data: token,
	})

}

func (u *UserController) Register(ctx echo.Context) error {
	var user structs.User
	ctx.Bind(&user)
	if err := validate.Struct(user); err != nil {
		verr, _ := err.(validator.ValidationErrors)
		return ctx.JSON(http.StatusConflict, structs.ApiResult{
			Error: utils.ExtractErrorMessages(verr),
		})
	}
	modelUser := model.NewUser()
	existUser, ok := modelUser.Find(user.Name, user.Password)
	if !ok {
		return ctx.JSON(http.StatusConflict, structs.ApiResult{
			Error: "Error in the DB",
		})
	}

	if existUser.Name != "" {
		return ctx.JSON(http.StatusConflict, structs.ApiResult{
			Error: "The User exist",
		})
	}

	registerUser, ok := modelUser.Register(user)

	if !ok {
		return ctx.JSON(http.StatusConflict, structs.ApiResult{
			Error: "Error register the user in the DB",
		})
	}

	token, err := utils.CreateJsonWebToken(registerUser)

	if err != nil {
		return ctx.JSON(http.StatusConflict, structs.ApiResult{
			Error: "Error creating the jwt",
		})
	}

	return ctx.JSON(http.StatusOK, structs.ApiResult{
		Data: token,
	})
}
