package entity

type Product struct {
	Id int `json:"id"`
	Name string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required,min=50"`
	ImageUrl string `json:"image_url" binding:"required"`
	Price int `json:"price" binding:"required,min=1000"`
	StoreId int `json:"store_id" binding:"required"`
}