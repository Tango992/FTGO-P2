package entity

type Transaction struct {
	ID          uint    `json:"transaction_id" extensions:"x-order=0"`
	UserID      uint    `json:"user_id" extensions:"x-order=1"`
	ProductID   uint    `json:"product_id" extensions:"x-order=2"`
	StoreID     uint    `json:"store_id" extensions:"x-order=3"`
	Quantity    uint    `json:"quantity" extensions:"x-order=4"`
	TotalAmount float32 `json:"total_amount" extensions:"x-order=5"`
}
