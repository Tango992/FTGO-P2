package handler

import (
	"net/http"
	"ungraded-5/entity"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	*DbHandler
}

func NewUserHandler(dbHandler *DbHandler) *UserHandler {
	return &UserHandler{
		DbHandler: dbHandler,
	}
}

func (u UserHandler) Register(c *gin.Context) {
	var user entity.User
	if err := c.ShouldBindJSON(&user); err != nil {
		WriteJson(&c, entity.Response{
			Code: http.StatusBadRequest,
			Message: "Failed to bind JSON",
			Data: nil,
		})
		return
	}

	if reflectErr := ValidateStruct(&user); reflectErr != nil {
		WriteJson(&c, *reflectErr)
		return
	}

	if hashErr := CreateHash(&user); hashErr != nil {
		WriteJson(&c, *hashErr)
		return
	}

	if dbErr := u.DbHandler.AddUserToDb(user); dbErr != nil {
		WriteJson(&c, *dbErr)
		return
	}

	WriteJson(&c, entity.Response{
		Code: http.StatusCreated,
		Message: "Registered successfully",
		Data: nil,
	})
}

func (u UserHandler) Login(c *gin.Context) {
	var credential entity.Credential
	if err := c.ShouldBindJSON(&credential); err != nil {
		WriteJson(&c, entity.Response{
			Code: http.StatusBadRequest,
			Message: "Failed to bind JSON",
			Data: nil,
		})
		return
	}

	hash, dbErr := u.DbHandler.FindHashInDb(&credential)
	if dbErr != nil {
		WriteJson(&c, *dbErr)
		return
	}

	if hashErr := HashMatched(hash, &credential); hashErr != nil {
		WriteJson(&c, *hashErr)
		return
	}

	WriteJson(&c, entity.Response{
		Code: http.StatusOK,
		Message: "Logged in",
		Data: nil,
	})
}
