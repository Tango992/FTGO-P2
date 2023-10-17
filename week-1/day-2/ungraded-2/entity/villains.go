package entity

type Villains struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Universe string `json:"universe"`
	ImageURL string `json:"imgurl"`
}
