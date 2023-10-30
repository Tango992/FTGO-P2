package helpers

import (
	"preview-week3/dto"
	"preview-week3/entity"
	"preview-week3/utils"

	"golang.org/x/crypto/bcrypt"
)

func CreateHash(data *dto.RegisterData) *utils.ErrResponse {
	hashed, err:= bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		hashErr := utils.ErrInternalServer
		hashErr.Details = err.Error()
		return &hashErr
	}

	data.Password = string(hashed)
	return nil
}

func PasswordMismatch(dbData entity.User, data dto.LoginData) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(dbData.Password), []byte(data.Password)); err != nil {
		return true
	}
	return false
}