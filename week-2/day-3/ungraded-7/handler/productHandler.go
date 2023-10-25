package handler

import (
	"ugc-7/dto"

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

func (ph ProductHandler) Hello(c *gin.Context) {
	WriteJson(&c, &dto.Response{
		Code: 200,
		Message: "Authenticated",
		Data: nil,
	})
}