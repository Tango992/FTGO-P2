package entity

import "time"

type Photo struct {
	ID        uint      `gorm:"primaryKey" json:"id,omitempty"`
	Title     string    `gorm:"not null" json:"title" binding:"required"`
	Caption   string    `json:"caption"`
	PhotoUrl  string    `gorm:"not null" json:"photo_url" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Comments  []Comment `json:"-"`
}