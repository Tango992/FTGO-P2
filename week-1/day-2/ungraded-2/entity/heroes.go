package entity

type Heroes struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Universe string `json:"universe"`
	Skill    string `json:"skill"`
	ImageURL string `json:"imgurl"`
}
