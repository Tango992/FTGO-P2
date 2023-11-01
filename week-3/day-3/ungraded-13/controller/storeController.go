package controller

import (
	"net/http"
	"ungraded-13/repository"

	"github.com/labstack/echo/v4"
)

type StoreController struct {
	repository.DbHandler
}

func NewStoreController(dbHandler repository.DbHandler) StoreController {
	return StoreController{
		DbHandler: dbHandler,
	}
}

func (sc StoreController) GetStores(c echo.Context) error {
	stores, dbErr := sc.DbHandler.FindAllStoresInDb()
	if dbErr != nil {
		return dbErr
	}
	return c.JSON(http.StatusOK,stores)
}

func (sc StoreController) GetStoreById(c echo.Context) error {
	return nil
}