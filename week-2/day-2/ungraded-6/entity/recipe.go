package entity

type Recipe struct {
	Id int `json:"id,omitempty"`
	Name string `json:"name" required:"true"`
	Description string `json:"description" required:"true"`
	Duration float32 `json:"duration" required:"true"`
	Rating float32 `json:"rating" required:"true"`
}