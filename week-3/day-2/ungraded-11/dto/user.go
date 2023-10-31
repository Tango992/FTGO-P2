package dto

type RegisterUser struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	DepositAmount float32 `json:"deposit_amount" validate:"required"`
}