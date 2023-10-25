package entity

type Store struct {
	Id int `json:"id,omitempty"`
	Email string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
	Name string `json:"store_name" binding:"required,min=6,max=15"`
	Type string `json:"store_type" binding:"required,oneof=silver gold platinum"`
}