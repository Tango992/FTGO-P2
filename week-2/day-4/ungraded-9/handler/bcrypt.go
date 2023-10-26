package handler

import (
	"ugc-9/dto"
	"ugc-9/entity"
	"ugc-9/utils"

	"golang.org/x/crypto/bcrypt"
)

func CreateHash(store *entity.Store) *utils.ErrResponse {
	hash, err := bcrypt.GenerateFromPassword([]byte(store.Password), bcrypt.DefaultCost)
	if err != nil {
		hashErr := utils.ErrInternalServer
		hashErr.Description = err.Error()
		return &hashErr
	}

	store.Password = string(hash)
	return nil
}

func HashMatched(hash string, credential *dto.Credential) *utils.ErrResponse {
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(credential.Password)); err != nil {
		hashErr := utils.ErrUnauthorized
		hashErr.Description = "Invalid credentials"
		return &hashErr
	}
	return nil
}