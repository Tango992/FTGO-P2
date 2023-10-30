package handler

import (
	"net/http"
	"os"
	"preview-week3/dto"
	"preview-week3/entity"
	"preview-week3/helpers"
	"preview-week3/utils"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
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
		Role: registerTemp.Role,
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

func (uc UserController) Login(c *gin.Context) {
	var loginData dto.LoginData

	if err := c.ShouldBindJSON(&loginData); err != nil {
		helpers.ErrJsonWriter(c, utils.ErrBadRequest, err.Error())
		return
	}

	dbUserData, dbErr := uc.DbHandler.FindUserInDb(loginData)
	if dbErr != nil {
		helpers.ErrJsonWriter(c, *dbErr, nil)
		return
	}

	if helpers.PasswordMismatch(dbUserData, loginData) {
		helpers.ErrJsonWriter(c, utils.ErrUnauthorized, "Invalid email or password")
		return
	}

	claims := jwt.MapClaims{
		"id": dbUserData.ID,
		"username": dbUserData.Username,
		"email": dbUserData.Email,
	}
	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenSigned, jwtErr := token.SignedString([]byte(os.Getenv("SECRET")))
	if jwtErr != nil {
		helpers.ErrJsonWriter(c, utils.ErrInternalServer, "Failed to sign jwt token")
	}

	dbUserData.Password = ""
	c.JSON(http.StatusOK, dto.Response{
		Message: "Logged in. Use the following token for authorization purposes",
		Data: tokenSigned,
	})
}