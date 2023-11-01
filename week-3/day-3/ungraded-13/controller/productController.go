package controller

import (
	"net/http"
	"ungraded-13/entity"
	"ungraded-13/dto"
	"ungraded-13/repository"

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

// Products      godoc
// @Summary      View all products
// @Tags         products
// @Produce      json
// @Success      200  {object}  []entity.Product
// @Failure      400  {object}  dto.Error
// @Failure      500  {object}  dto.Error
// @Router       /products [get]
func (pc ProductController) GetProducts(c echo.Context) error {
	products, err := pc.DbHandler.GetAllProducts()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, products)
}

// Transaction   godoc
// @Summary      Establish a transaction
// @Tags         transactions
// @Accept       json
// @Produce      json
// @Param        Authorization header string true "JWT Token"
// @Param        request body dto.RequestTransaction  true  "Transaction data"
// @Success      201  {object}  dto.TransactionResponse
// @Failure      400  {object}  dto.Error
// @Failure      401  {object}  dto.Error
// @Failure      404  {object}  dto.Error
// @Failure      500  {object}  dto.Error
// @Router       /transactions [post]
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
		ProductID: requestTransactionTemp.ProductID,
		Quantity: requestTransactionTemp.Quantity,
		StoreID: requestTransactionTemp.StoreID,
	}

	if dbErr := pc.DbHandler.EstablishTransactions(&requestTransaction); dbErr != nil {
		return dbErr
	}
	
	return c.JSON(http.StatusCreated, dto.Response{
		Message: "Transaction established",
		Data: requestTransaction,
	})
}