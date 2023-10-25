package dto

type Response struct {
	Code int `json:"code"`
	Message string `json:"message"`
	Data any `json:"data"`
}

type Credential struct {
	Email string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}