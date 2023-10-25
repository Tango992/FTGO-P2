package handler

import (
	"net/http"
	"ugc-8/dto"
	"ugc-8/entity"

	"golang.org/x/crypto/bcrypt"
)

func CreateHash(store *entity.Store) *dto.Response {
	hash, err := bcrypt.GenerateFromPassword([]byte(store.Password), bcrypt.DefaultCost)
	if err != nil {
		return &dto.Response{
			Code: http.StatusInternalServerError,
			Message: err.Error(),
			Data: nil,
		}
	}

	store.Password = string(hash)
	return nil
}

func HashMatched(hash string, credential *dto.Credential) *dto.Response {
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(credential.Password)); err != nil {
		return &dto.Response{
			Code: http.StatusUnauthorized,
			Message: "Invalid credentials",
			Data: nil,
		}
	}
	return nil
}