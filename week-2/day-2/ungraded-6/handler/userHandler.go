package handler

import (
	"net/http"
	"os"
	"time"
	"ungraded-6/entity"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
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

	user, dbErr := u.DbHandler.FindUserInDb(&credential)
	if dbErr != nil {
		WriteJson(&c, *dbErr)
		return
	}

	if hashErr := HashMatched(user.Password, &credential); hashErr != nil {
		WriteJson(&c, *hashErr)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": user.Id,
		"name": user.Name,
		"role": user.Role,
		"exp": time.Now().Add(time.Hour).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		WriteJson(&c, entity.Response{
			Code: http.StatusInternalServerError,
			Message: "Failed to sign token",
			Data: nil,
		})
	}

	// Send same site cookie
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600, "", "", false, true)

	WriteJson(&c, entity.Response{
		Code: http.StatusOK,
		Message: "Logged in",
		Data: nil,
	})
}
