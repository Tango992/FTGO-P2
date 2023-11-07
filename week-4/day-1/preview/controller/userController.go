package controller

import (
	"net/http"
	"preview-w4/models"
	"preview-w4/repository"
	"preview-w4/utils"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	repository.DbHandler
}

func NewUserController(dbHandler repository.DbHandler) UserController {
	return UserController{
		DbHandler: dbHandler,
	}
}

func (uc UserController) Register(c echo.Context) error {
	var registerData models.User
	if err := c.Bind(&registerData); err != nil {
		return echo.NewHTTPError(utils.ErrBadRequest.Details(err.Error()))
	}

	if err := c.Validate(&registerData); err != nil {
		return echo.NewHTTPError(utils.ErrBadRequest.Details(err.Error()))
	}

	if err := uc.DbHandler.AddUserIntoDb(&registerData); err != nil {
		return err
	}
	
	return c.JSON(http.StatusCreated, models.Response{
		Message: "Registered",
		Data: registerData,
	})
}