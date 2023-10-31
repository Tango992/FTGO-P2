package controller

import (
	"fmt"
	"net/http"
	"ungraded-11/entity"
	"ungraded-11/dto"
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
	claimsTemp := c.Get("user")
	claims := claimsTemp.(map[string]any)

	var requestTransactionTemp dto.RequestTransaction
	if err := c.Bind(&requestTransactionTemp); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(&requestTransactionTemp); err != nil {
		return err
	}

	requestTransaction := entity.Transaction{
		UserID: uint(claims["id"].(float64)),
		ProductID: requestTransactionTemp.ProductId,
		Quantity: requestTransactionTemp.Quantity,
	}

	if dbErr := pc.DbHandler.EstablishTransactions(&requestTransaction); dbErr != nil {
		return dbErr
	}
	
	return c.JSON(http.StatusCreated, echo.Map{
		"message": "Transaction established",
		"data": requestTransaction,
	})
}