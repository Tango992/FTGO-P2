package handler

import (
	"net/http"
	"strconv"
	"ugc-8/dto"
	"ugc-8/entity"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	*DbHandler
}

func NewProductHandler(dbhandler *DbHandler) *ProductHandler {
	return &ProductHandler{
		DbHandler: dbhandler,
	}
}

func (ph ProductHandler) PostProduct(c *gin.Context) {
	var product entity.Product

	// Get user id from cookie
	user, _ := c.Get("user")
	id := uint(user.(gin.H)["id"].(float64))
	product.StoreID = id
	
	if err := c.ShouldBindJSON(&product); err != nil {
		WriteJson(&c, &dto.Response{
			Code: http.StatusBadRequest,
			Message: err.Error(),
			Data: nil,
		})
		return
	}

	if dbErr := ph.DbHandler.AddProductToDb(&product); dbErr != nil {
		WriteJson(&c, dbErr)
		return
	}

	WriteJson(&c, &dto.Response{
		Code: http.StatusCreated,
		Message: "Product posted",
		Data: product,
	})
}

func (ph ProductHandler) GetAllProducts(c *gin.Context) {
	products, dbErr := ph.DbHandler.FindAllProductFromDb()
	if dbErr != nil {
		WriteJson(&c, dbErr)
		return
	}

	WriteJson(&c, &dto.Response{
		Code: http.StatusOK,
		Message: "Get all products",
		Data: products,
	})
}

func (ph ProductHandler) GetProductById(c *gin.Context) {
	var product entity.Product
	
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		WriteJson(&c, &dto.Response{
			Code: http.StatusBadRequest,
			Message: "Invalid syntax",
			Data: nil,
		})
		return
	}
	product.ID = uint(id)

	if dbErr := ph.DbHandler.FindProductInDb(&product); dbErr != nil {
		WriteJson(&c, dbErr)
		return
	}

	WriteJson(&c, &dto.Response{
		Code: http.StatusOK,
		Message: "Get product by ID",
		Data: product,
	})
}

func (ph ProductHandler) DeleteProductById(c *gin.Context) {
	var product entity.Product

	// Get user id from cookie
	user, _ := c.Get("user")
	storeId := uint(user.(gin.H)["id"].(float64))
	product.StoreID = storeId
	
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		WriteJson(&c, &dto.Response{
			Code: http.StatusBadRequest,
			Message: "Invalid syntax",
			Data: nil,
		})
		return
	}
	product.ID = uint(id)

	if dbErr := ph.DbHandler.DeleteProductInDb(&product); dbErr != nil {
		WriteJson(&c, dbErr)
		return
	}

	WriteJson(&c, &dto.Response{
		Code: http.StatusOK,
		Message: "Delete product by ID",
		Data: nil,
	})
}

func (ph ProductHandler) UpdateProductById(c *gin.Context) {
	var product entity.Product

	// Get user id from cookie
	user, _ := c.Get("user")
	storeId := uint(user.(gin.H)["id"].(float64))
	product.StoreID = storeId
	
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		WriteJson(&c, &dto.Response{
			Code: http.StatusBadRequest,
			Message: "Invalid syntax",
			Data: nil,
		})
		return
	}
	product.ID = uint(id)

	if err := c.ShouldBindJSON(&product); err != nil {
		WriteJson(&c, &dto.Response{
			Code: http.StatusBadRequest,
			Message: err.Error(),
			Data: nil,
		})
		return
	}

	if dbErr := ph.DbHandler.UpdateProductInDb(&product); dbErr != nil {
		WriteJson(&c, dbErr)
		return
	}

	WriteJson(&c, &dto.Response{
		Code: http.StatusOK,
		Message: "Product updated",
		Data: product,
	})
}