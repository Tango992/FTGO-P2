package models

type Player struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"not null"`
	TeamName string `gorm:"not null"`
	Ranking  uint   `gorm:"not null"`
	Score    uint   `gorm:"not null"`
}
