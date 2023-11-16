package controllers

import (
	model "main/src/Model"
	structs "main/src/Structs"
	"main/src/utils"
	"net/http"

	"github.com/labstack/echo"
)

func Login(ctx echo.Context) error {
	var user structs.User
	ctx.Bind(&user)
	modelUser := model.NewUser()
	existUser, ok := modelUser.FindByName(user.Name)

	if !ok {
		return ctx.JSON(http.StatusConflict, map[string]string{
			"error": "Error in the DB",
		})
	}

	if existUser.Name == "" {
		return ctx.JSON(http.StatusConflict, map[string]string{
			"error": "The User not exist",
		})
	}

	token, err := utils.CreateJsonWebToken(existUser)

	if err != nil {
		return ctx.JSON(http.StatusConflict, map[string]string{
			"error": "Error creating the jwt",
		})
	}

	return ctx.JSON(http.StatusConflict, map[string]string{
		"message": token,
	})

}

func Register(ctx echo.Context) error {
	var user structs.User
	ctx.Bind(&user)
	modelUser := model.NewUser()
	existUser, ok := modelUser.FindByName(user.Name)
	if !ok {
		return ctx.JSON(http.StatusConflict, map[string]string{
			"error": "Error in the DB",
		})
	}

	if existUser.Name != "" {
		return ctx.JSON(http.StatusConflict, map[string]string{
			"error": "The User exist",
		})
	}

	registerUser, ok := modelUser.Register(user)

	if !ok {
		return ctx.JSON(http.StatusConflict, map[string]string{
			"error": "Error register the user in the DB",
		})
	}

	token, err := utils.CreateJsonWebToken(registerUser)

	if err != nil {
		return ctx.JSON(http.StatusConflict, map[string]string{
			"error": "Error creating the jwt",
		})
	}

	return ctx.JSON(http.StatusConflict, map[string]string{
		"message": token,
	})
}
