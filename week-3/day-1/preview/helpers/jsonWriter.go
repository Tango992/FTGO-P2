package helpers

import (
	"preview-week3/utils"

	"github.com/gin-gonic/gin"
)

func ErrJsonWriter(c *gin.Context, contract utils.ErrResponse, data any) {
	if contract.Details == nil && data == nil {
		contract.Details = struct{}{}
		c.JSON(contract.Code, contract)
		return
	}

	if data == nil {
		c.JSON(contract.Code, contract)
		return
	}

	contract.Details = data
	c.JSON(contract.Code, contract)
}