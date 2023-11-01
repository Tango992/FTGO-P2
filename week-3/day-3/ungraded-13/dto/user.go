package dto

type RegisterUser struct {
	Username      string  `json:"username" validate:"required" extensions:"x-order=0"`
	Password      string  `json:"password" validate:"required" extensions:"x-order=1"`
	DepositAmount float32 `json:"deposit_amount" validate:"required" extensions:"x-order=2"`
}

type Login struct {
	Username string `json:"username" validate:"required" extensions:"x-order=0"`
	Password string `json:"password" validate:"required" extensions:"x-order=1"`
}
