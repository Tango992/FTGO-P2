package handler

import (
	"net/http"
	"os"
	"time"
	"ugc-9/dto"
	"ugc-9/entity"
	"ugc-9/utils"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type StoreHandler struct {
	*DbHandler
}

func NewStoreHandler(dbHandler *DbHandler) *StoreHandler {
	return &StoreHandler{
		DbHandler: dbHandler,
	}
}

func (u StoreHandler) Register(c *gin.Context) {
	var store entity.Store
	if err := c.ShouldBindJSON(&store); err != nil {
		WriteJsonErr(c, &utils.ErrBadRequest, err.Error())
		return
	}

	if hashErr := CreateHash(&store); hashErr != nil {
		WriteJsonErr(c, hashErr)
		return
	}

	if dbErr := u.DbHandler.AddStoreToDb(store); dbErr != nil {
		WriteJsonErr(c, dbErr)
		return
	}

	WriteJson(&c, &dto.Response{
		Code: http.StatusCreated,
		Message: "Registered successfully",
		Data: nil,
	})
}

func (u StoreHandler) Login(c *gin.Context) {
	var credential dto.Credential
	if err := c.ShouldBindJSON(&credential); err != nil {
		WriteJsonErr(c, &utils.ErrBadRequest, err.Error())
		return
	}

	store, dbErr := u.DbHandler.FindStoreInDb(&credential)
	if dbErr != nil {
		WriteJsonErr(c, dbErr)
		return
	}

	if hashErr := HashMatched(store.Password, &credential); hashErr != nil {
		WriteJsonErr(c, hashErr)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": store.ID,
		"name": store.Name,
		"email": store.Email,
		"role": store.Type,
		"exp": time.Now().Add(time.Hour).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		WriteJson(&c, &dto.Response{
			Code: http.StatusInternalServerError,
			Message: "Failed to sign token",
			Data: nil,
		})
		return
	}

	// Send same site cookie
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600, "", "", true, true)

	WriteJson(&c, &dto.Response{
		Code: http.StatusOK,
		Message: "Logged in",
		Data: nil,
	})
}
