package repository

import (
	"preview-w4/models"
	"preview-w4/utils"

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

func (db DbHandler) AddUserIntoDb(data *models.User) error {
	if err := db.Create(data).Error; err != nil {
		return echo.NewHTTPError(utils.ErrInternalServer.Details(err.Error()))
	}
	return nil
}