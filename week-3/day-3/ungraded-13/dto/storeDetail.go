package dto

type StoreWithSales struct {
	ID         uint    `json:"store_id" extensions:"x-order=0"`
	Name       string  `json:"name" extensions:"x-order=1"`
	Rating     float32 `json:"rating,omitempty" extensions:"x-order=2"`
	Address    string  `json:"address" extensions:"x-order=3"`
	Lat        float64 `json:"latitude,omitempty" extensions:"x-order=4"`
	Long       float64 `json:"longitude,omitempty" extensions:"x-order=5"`
	TotalSales float64 `json:"total_sales" extensions:"x-order=6"`
}

type StoreDetail struct {
	StoreWithSales `json:"store"`
	Weather        any `json:"weather"`
}
