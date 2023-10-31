package dto

type RequestTransaction struct {
	ProductId uint `json:"product_id" validate:"required"`
	Quantity uint `json:"quantity" validate:"required"`
}