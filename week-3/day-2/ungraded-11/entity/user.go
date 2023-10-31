package entity

type User struct {
	ID            uint    `json:"id"`
	Username      string  `json:"username"`
	Password      string  `json:"password,omitempty"`
	DepositAmount float32 `json:"deposit_amount"`
}
