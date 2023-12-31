package dto

type RegisterData struct {
	Username string `gorm:"not null;unique" json:"username" binding:"required"`
	Email    string `gorm:"not null;unique" json:"email" binding:"required,email"`
	Password string `gorm:"not null" json:"password" binding:"required,min=6"`
	Age      uint   `gorm:"not null" json:"age" binding:"required,min=8"`
	Role     uint   `gorm:"not null" json:"role" binding:"required,min=0,max=1"`
}

type LoginData struct {
	Email    string `gorm:"not null;unique" json:"email" binding:"required,email"`
	Password string `gorm:"not null" json:"password" binding:"required,min=6"`
}
