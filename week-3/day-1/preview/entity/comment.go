package entity

import "time"

type Comment struct {
	ID        uint      `gorm:"primaryKey" json:"id,omitempty"`
	UserID    uint      `gorm:"not null" json:"user_id"`
	PhotoID   uint      `gorm:"not null" json:"photo_id"`
	Message   uint      `gorm:"not null" json:"message" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
