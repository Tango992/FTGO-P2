package models

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"not null" validate:"required"`
	Password string `gorm:"not null" validate:"required"`
}
