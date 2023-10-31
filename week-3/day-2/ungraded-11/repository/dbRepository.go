package repository

import (
	"errors"
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

func (db DbHandler) FindUserInDb(username string) (entity.User, error){
	var user entity.User

	res := db.Where("username = ?", username).First(&user)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return entity.User{}, echo.NewHTTPError(http.StatusUnauthorized, "Invalid credentials")
		}
		return entity.User{}, echo.NewHTTPError(http.StatusInternalServerError, res.Error.Error())
	}

	return user, nil
}

func (db DbHandler) GetAllProducts() ([]entity.Product, error) {
	products := []entity.Product{}

	res := db.Find(&products)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return []entity.Product{}, echo.NewHTTPError(http.StatusUnauthorized, "Invalid credentials")
		}
		return []entity.Product{}, echo.NewHTTPError(http.StatusInternalServerError, res.Error.Error())
	}
	return products, nil
}