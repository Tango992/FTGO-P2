package entity

type PostCrimeReport struct {
	ID int `json:"id"`
	Hero_id int `json:"hero_id"`
	Villain_id int `json:"villain_id"`
	Description string `json:"description"`
	Date string `json:"date"`
}

type GetCrimeReport struct {
	ID int `json:"id"`
	Hero `json:"hero"`
	Villain `json:"villain"`
	Description string `json:"description"`
	Date string `json:"date"`
}