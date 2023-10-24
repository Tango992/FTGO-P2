package entity

type User struct {
	Id         int    `json:"id,omitempty"`
	Email      string `json:"email" required:"true"`
	Password   string `json:"password" required:"true" minLen:"8" maxLen:"100"`
	Name       string `json:"name" required:"true" minLen:"6" maxLen:"15"`
	Age        int    `json:"age" required:"true" min:"17" max:"90"`
	Occupation string `json:"occupation" required:"true"`
	Role       string `json:"role" required:"true"`
}

type Credential struct {
	Email    string `json:"email" required:"true"`
	Password string `json:"password" required:"true"`
}
