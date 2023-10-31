package controller

import (
	"net/http"
	"os"
	"time"
	"ungraded-11/dto"
	"ungraded-11/helpers"
	"ungraded-11/repository"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	repository.DbHandler
}

func NewUserController(dbHandler repository.DbHandler) UserController {
	return UserController{
		DbHandler: dbHandler,
	}
}

// Register      godoc
// @Summary      Register an account
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        request body dto.RegisterUser  true  "Registration data"
// @Success      201  {object}  dto.RegisterResponse
// @Failure      400  {object}  dto.Error
// @Failure      500  {object}  dto.Error
// @Router       /users/register [post]
func (uc UserController) Register(c echo.Context) error {
	var registerData dto.RegisterUser

	if err := c.Bind(&registerData); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(&registerData); err != nil {
		return err
	}

	if err := helpers.CreateHashPassword(&registerData); err != nil {
		return err
	}

	user, dbErr := uc.DbHandler.AddUserIntoDb(registerData)
	if dbErr != nil {
		return dbErr
	}
	
	return c.JSON(http.StatusCreated, user)
}

// Login         godoc
// @Summary      Log in an existing account
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        request body dto.Login  true  "Log in data"
// @Success      200  {object}  dto.LoginResponse
// @Failure      400  {object}  dto.Error
// @Failure      401  {object}  dto.Error
// @Failure      404  {object}  dto.Error
// @Failure      500  {object}  dto.Error
// @Router       /users/login [post]
func (uc UserController) Login(c echo.Context) error {
	var loginData dto.Login

	if err := c.Bind(&loginData); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(&loginData); err != nil {
		return err
	}

	dbData, dbErr := uc.DbHandler.FindUserInDb(loginData.Username)
	if dbErr != nil {
		return dbErr
	}

	if helpers.PasswordDoesNotMatch(dbData, loginData) {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid username / passowrd")
	}

	claims := jwt.MapClaims{
		"exp": time.Now().Add(3*time.Hour).Unix(),
		"id": dbData.ID,
		"username": dbData.Username,
	}
	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	
	dbData.Password = ""
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Logged in!",
		"token": tokenString,
	})
}