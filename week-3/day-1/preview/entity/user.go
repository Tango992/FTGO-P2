package entity

import "time"

type User struct {
	ID          uint        `gorm:"primaryKey" json:"id,omitempty"`
	Username    string      `gorm:"not null;unique" json:"username" binding:"required"`
	Email       string      `gorm:"not null;unique" json:"email" binding:"required,email"`
	Password    string      `gorm:"not null" json:"password,omitempty" binding:"required,min=6"`
	Age         uint        `gorm:"not null" json:"age" binding:"required,min=8"`
	Role        uint        `gorm:"not null" json:"role" binding:"required,min=0,max=1"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
	Photos      []Photo     `json:"-"`
	Comments    []Comment   `json:"-"`
	SocialMedia SocialMedia `json:"-"`
}
