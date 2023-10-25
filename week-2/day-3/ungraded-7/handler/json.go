package handler

import (
	"ugc-7/dto"
	"github.com/gin-gonic/gin"
)

func WriteJson(c **gin.Context, data *dto.Response) {
	if data.Data == nil {
		(*c).JSON(data.Code, gin.H{
			"code": data.Code,
			"message": data.Message,
		})
		return
	}

	(*c).JSON(data.Code, dto.Response{
		Code: data.Code,
		Message: data.Message,
		Data: data.Data,
	})
}