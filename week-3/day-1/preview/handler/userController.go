package handler

import (
	"net/http"
	"preview-week3/dto"
	"preview-week3/entity"
	"preview-week3/helpers"
	"preview-week3/utils"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	DbHandler
}

func NewUserController(dbHandler DbHandler) UserController {
	return UserController{
		DbHandler: dbHandler,
	}
}

func (uc UserController) Register(c *gin.Context) {
	var registerTemp dto.RegisterData

	if err := c.ShouldBindJSON(&registerTemp); err != nil {
		helpers.ErrJsonWriter(c, utils.ErrBadRequest, err.Error())
		return
	}

	if hashErr := helpers.CreateHash(&registerTemp); hashErr != nil {
		helpers.ErrJsonWriter(c, *hashErr, nil)
		return
	}

	registerData := entity.User{
		Username: registerTemp.Username,
		Email: registerTemp.Email,
		Password: registerTemp.Password,
		Age: registerTemp.Age,
	}

	if dbErr := uc.DbHandler.AddUserIntoDb(&registerData); dbErr != nil {
		helpers.ErrJsonWriter(c, *dbErr, nil)
		return
	}
	
	c.JSON(http.StatusCreated, dto.Response{
		Message: "Registered successfully",
		Data: registerData,
	})
}