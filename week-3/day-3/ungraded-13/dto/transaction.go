package dto

type RequestTransaction struct {
	ProductID uint `json:"product_id" validate:"required"`
	Quantity  uint `json:"quantity" validate:"required"`
	StoreID   uint `json:"store_id" validate:"required"`
}
