package models

type User struct {
	UserID    int    `json:"user_Id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Role      string `json:"role"`
	Gender    string `json:"gender"`
	Address   string `json:"address"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
