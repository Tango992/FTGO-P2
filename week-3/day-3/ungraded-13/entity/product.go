package entity

type Product struct {
	ID    uint    `json:"product_id" extensions:"x-order=0"`
	Name  string  `json:"name" extensions:"x-order=1"`
	Stock uint    `json:"stock" extensions:"x-order=2"`
	Price float32 `json:"price" extensions:"x-order=3"`
}
