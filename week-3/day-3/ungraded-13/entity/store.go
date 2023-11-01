package entity

type Store struct {
	ID      uint    `json:"store_id"`
	Name    string  `json:"name"`
	Rating  float32 `json:"rating,omitempty"`
	Address string  `json:"address"`
	Lat     float64 `json:"latitude,omitempty"`
	Long    float64 `json:"longitude,omitempty"`
}
