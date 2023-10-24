package handler

import (
	"net/http"
	"ungraded-6/entity"

	"golang.org/x/crypto/bcrypt"
)

func CreateHash(user *entity.User) *entity.Response {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return &entity.Response{
			Code: http.StatusInternalServerError,
			Message: err.Error(),
			Data: nil,
		}
	}

	user.Password = string(hash)
	return nil
}

func HashMatched(hash string, credential *entity.Credential) *entity.Response {
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(credential.Password)); err != nil {
		return &entity.Response{
			Code: http.StatusUnauthorized,
			Message: "Invalid credentials",
			Data: nil,
		}
	}
	return nil
}