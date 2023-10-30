package handler

import (
	"errors"
	"preview-week3/dto"
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

func (db DbHandler) FindUserInDb(data dto.LoginData) (entity.User, *utils.ErrResponse) {
	var user entity.User

	res := db.Where("email = ?", data.Email).First(&user)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			resErr := utils.ErrNotFound
			resErr.Details = res.Error.Error()
			return entity.User{}, &resErr
		}

		resErr := utils.ErrInternalServer
		resErr.Details = res.Error.Error()
		return entity.User{}, &resErr
	}
	
	return user, nil
}