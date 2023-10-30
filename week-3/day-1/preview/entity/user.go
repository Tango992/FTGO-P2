package entity

import "time"

type User struct {
	ID          uint        `gorm:"primaryKey" json:"id,omitempty"`
	Username    string      `gorm:"not null;unique" json:"username" binding:"required"`
	Email       string      `gorm:"not null;unique" json:"email" binding:"required,email"`
	Password    string      `gorm:"not null" json:"password" binding:"required,min=6"`
	Age         uint        `gorm:"not null" json:"age" binding:"required,min=8"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
	Comments    []Comment   `json:"-"`
	SocialMedia SocialMedia `json:"-"`
}
