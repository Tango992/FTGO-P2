package handler

import (
	"net/http"
	"ungraded-11/dto"
	"ungraded-11/entity"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type DbHandler struct {
	*gorm.DB
}

func NewDbHandler(db *gorm.DB) DbHandler {
	return DbHandler{
		DB: db,
	}
}

func (db DbHandler) AddUserIntoDb(data dto.RegisterUser) (entity.User, error) {
	user := entity.User{
		Username: data.Username,
		Password: data.Password,
		DepositAmount: data.DepositAmount,
	}

	res := db.Create(&user)
	if res.Error != nil {
		return entity.User{}, echo.NewHTTPError(http.StatusInternalServerError, res.Error.Error())
	}

	user.Password = ""
	return user, nil
}