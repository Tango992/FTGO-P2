package entity

type Branch struct {
	Branch_id int `json:"branch_id,omitempty"`
	Name string `json:"name"`
	Location string `json:"location"`
}