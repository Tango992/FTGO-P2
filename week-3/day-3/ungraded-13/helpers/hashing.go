package helpers

import (
	"net/http"
	"ungraded-13/entity"
	"ungraded-13/dto"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func CreateHashPassword(data *dto.RegisterUser) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create hash")
	}

	data.Password = string(hashedPassword)
	return nil
}

func PasswordDoesNotMatch(dbData entity.User, loginData dto.Login) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(dbData.Password), []byte(loginData.Password)); err != nil {
		return true
	}
	return false
}