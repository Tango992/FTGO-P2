package entity

type Store struct {
	ID       uint   `gorm:"primaryKey" json:"id,omitempty"`
	Email    string `gorm:"unique;not null" json:"email" binding:"required,email"`
	Password string `gorm:"not null" json:"password" binding:"required,min=8"`
	Name     string `gorm:"not null" json:"store_name" binding:"required,min=6,max=15"`
	Type     string `gorm:"not null" json:"store_type" binding:"required,oneof=silver gold platinum"`
	Products []Product
}
