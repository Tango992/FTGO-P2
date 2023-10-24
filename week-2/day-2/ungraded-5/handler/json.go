package handler

import (
	"ungraded-5/entity"
	"github.com/gin-gonic/gin"
)

func WriteJson(c **gin.Context, data entity.Response) {
	if data.Data == nil {
		(*c).JSON(data.Code, gin.H{
			"code": data.Code,
			"message": data.Message,
		})
		return
	}

	(*c).JSON(data.Code, entity.Response{
		Code: data.Code,
		Message: data.Message,
		Data: data.Data,
	})
}