package entity

import "time"

type SocialMedia struct {
	ID             uint      `gorm:"primaryKey" json:"id,omitempty"`
	Name           string    `gorm:"not null" json:"name" binding:"required"`
	SocialMediaUrl string    `gorm:"not null" binding:"required"`
	UserID         uint      `gorm:"not null" json:"user_id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}