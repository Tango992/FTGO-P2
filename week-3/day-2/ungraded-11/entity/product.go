package entity

type Product struct {
	ID    uint    `json:"product_id"`
	Name  string  `json:"name"`
	Stock uint    `json:"stock"`
	Price float32 `json:"price"`
}
