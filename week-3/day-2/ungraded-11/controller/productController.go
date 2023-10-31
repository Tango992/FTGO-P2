package controller

import (
	"fmt"
	"net/http"
	"ungraded-11/repository"

	"github.com/labstack/echo/v4"
)

type ProductController struct {
	repository.DbHandler
}

func NewProductHandler(dbHandler repository.DbHandler) ProductController {
	return ProductController{
		DbHandler: dbHandler,
	}
}

func (pc ProductController) GetProducts(c echo.Context) error {
	products, err := pc.DbHandler.GetAllProducts()
	if err != nil {
		return err
	}
	
	user := c.Get("user")
	fmt.Println(user)
	
	return c.JSON(http.StatusOK, products)
}

func (pc ProductController) PostTransaction(c echo.Context) error {
	return c.JSON(http.StatusCreated, echo.Map{
		"message": "Transaction",
	})
}