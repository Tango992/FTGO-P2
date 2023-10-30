package dto

type PhotoData struct {
	Title     string    `gorm:"not null" json:"title" binding:"required"`
	Caption   string    `json:"caption"`
	PhotoUrl  string    `gorm:"not null" json:"photo_url" binding:"required"`
}