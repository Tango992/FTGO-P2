package entity

type Transaction struct {
	ID          uint    `json:"transaction_id"`
	UserID      uint    `json:"user_id"`
	ProductID   uint    `json:"product_id"`
	Quantity    uint    `json:"quantity"`
	TotalAmount float32 `json:"total_amount"`
}
