package models

import "github.com/uptrace/bun"

type User struct {
	bun.BaseModel `bun:"table:users"`

	UserID    int64  `bun:",pk,autoincrement" json:"user_id"`
	FirstName string `bun:"firstname" json:"firstname"`
	LastName  string `bun:"lastname" json:"lastname"`
	Email     string `bun:"email" json:"email"`
	Password  string `bun:"password" json:"password"`
	Role      string `bun:"role" json:"role"`
	Gender    string `bun:"gender" json:"gender"`
	Address   string `bun:"address" json:"address"`
}

type LoginRequest struct {
	Email    string `bun:"email" json:"email"`
	Password string `bun:"password" json:"password"`
}
