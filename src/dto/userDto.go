package dto

type UserBody struct {
	Name     string `json:"name"`
	Number   string `json:"number"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
