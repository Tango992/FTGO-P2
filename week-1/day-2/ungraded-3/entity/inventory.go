package entity

type Inventory struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Stock       int    `json:"stock"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type PostInventory struct {
	Name        string `json:"name"`
	Stock       int    `json:"stock"`
	Description string `json:"description"`
	Status_id   int    `json:"status_id"`
}

type PutInventory struct {
	Stock       int    `json:"stock"`
	Status_id   int    `json:"status_id"`
}