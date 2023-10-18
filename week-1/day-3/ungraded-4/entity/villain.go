package entity

type Villain struct {
	V_ID       int    `json:"id"`
	V_Name     string `json:"name"`
	V_Universe string `json:"universe"`
	V_ImageURL string `json:"imgurl"`
}
