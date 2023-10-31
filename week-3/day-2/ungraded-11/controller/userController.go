package controller

import (
	"net/http"
	"ungraded-11/dto"
	"ungraded-11/handler"
	"ungraded-11/helpers"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	handler.DbHandler
}

func NewUserController(dbHandler handler.DbHandler) UserController {
	return UserController{
		DbHandler: dbHandler,
	}
}

func (uc UserController) Register(c echo.Context) error {
	var registerData dto.RegisterUser

	if err := c.Bind(&registerData); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(&registerData); err != nil {
		return err
	}

	if err := helpers.CreateHashPassword(&registerData); err != nil {
		return err
	}

	user, dbErr := uc.DbHandler.AddUserIntoDb(registerData)
	if dbErr != nil {
		return dbErr
	}
	
	return c.JSON(http.StatusCreated, user)
}

func (uc UserController) Login(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Logged in!",
	})
}