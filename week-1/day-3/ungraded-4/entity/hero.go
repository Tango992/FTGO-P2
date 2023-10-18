package entity

type Hero struct {
	H_ID       int    `json:"id"`
	H_Name     string `json:"name"`
	H_Universe string `json:"universe"`
	H_Skill    string `json:"skill"`
	H_ImageURL string `json:"imgurl"`
}
