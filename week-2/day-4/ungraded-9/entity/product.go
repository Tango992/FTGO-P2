package entity

type Product struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Name        string `gorm:"not null" json:"name" binding:"required"`
	Description string `gorm:"not null" json:"description" binding:"required,min=50"`
	ImageUrl    string `gorm:"not null" json:"image_url" binding:"required"`
	Price       int    `gorm:"not null" json:"price" binding:"required,min=1000"`
	StoreID     uint   `json:"store_id"`
}
