package controller

import (
	"net/http"
	"preview-w4/models"
	"preview-w4/repository"

	"github.com/labstack/echo/v4"
)

type PlayerController struct {
	repository.DbHandler
}

func NewPlayerController(dbHandler repository.DbHandler) PlayerController {
	return PlayerController{
		DbHandler: dbHandler,
	}
}

func (pc PlayerController) GetPlayers(c echo.Context) error {
	return c.JSON(http.StatusOK, models.Response{
		Message: "Get players",
	})
}

func (pc PlayerController) PostPlayer(c echo.Context) error {
	return c.JSON(http.StatusCreated, models.Response{
		Message: "Post players",
	})
}

func (pc PlayerController) PutPlayer(c echo.Context) error {
	return c.JSON(http.StatusOK, models.Response{
		Message: "Put players",
	})
}