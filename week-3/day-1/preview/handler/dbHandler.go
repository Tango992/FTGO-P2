package handler

import (
	"preview-week3/entity"
	"preview-week3/utils"

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

func (db DbHandler) AddUserIntoDb(data *entity.User) *utils.ErrResponse {
	res := db.Create(data)
	if res.Error != nil {
		resErr := utils.ErrInternalServer
		resErr.Details = res.Error.Error()
		return &resErr
	}

	data.Password = ""
	return nil
}