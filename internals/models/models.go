package models

import (
	"github.com/golang-jwt/jwt"
	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:users"`

	UserID    int64  `bun:",pk,autoincrement" json:"user_id"`
	FirstName string `bun:"firstname" json:"firstname"`
	LastName  string `bun:"lastname" json:"lastname"`
	Email     string `bun:"email" json:"email"`
	Password  string `bun:"password" json:"pass_word"`
	Role      string `bun:"role" json:"user_role"`
	Gender    string `bun:"gender" json:"gender"`
	Address   string `bun:"address" json:"user_address"`
}

type LoginRequest struct {
	Email    string `bun:"email" json:"email"`
	Password string `bun:"password" json:"password"`
}

type Claims struct {
	ID    int64
	Email string
	Role  string
	jwt.StandardClaims
}
