package handler

import (
	"ugc-9/dto"
	"ugc-9/utils"

	"github.com/gin-gonic/gin"
)

func WriteJsonErr(c **gin.Context, err *utils.ErrResponse, data ...any) {
	// If variadic arguments doesn't get called
	if len(data) < 1 {
		if err.Description == nil {
			err.Description = struct{}{}
			(*c).JSON(err.Code, err)
			return
		}
		
		(*c).JSON(err.Code, err)
		return
	}

	// If variadic arguments got called
	if len(data) == 1 {
		err.Description = data[0]
		(*c).JSON(err.Code, err)
		return
	}
	
	// If variadic arguments called multiple times
	err.Description = data
	(*c).JSON(err.Code, err)
}


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