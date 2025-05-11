package models

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:users"`

	UserID    int64  `bun:"user_id,pk,autoincrement" json:"user_id"`
	FirstName string `bun:"firstname,notnull" json:"firstname"`
	LastName  string `bun:"lastname,notnull" json:"lastname"`
	Email     string `bun:"email,unique,notnull" json:"email"`
	Password  string `bun:"pass_word,notnull" json:"pass_word"`
	Role      string `bun:"user_role,notnull,default:'patient'" json:"user_role"`
	Gender    string `bun:"gender" json:"gender"`
	Address   string `bun:"user_address" json:"user_address"`
}

type Appointment struct {
	bun.BaseModel `bun:"table:appointments"`

	AppointmentID int64     `bun:"appointment_id,pk,autoincrement" json:"appointment_id"`
	PatientID     int64     `bun:"patient_id,notnull" json:"patient_id"`
	StaffID       *int64    `bun:"staff_id,nullzero" json:"staff_id,omitempty"`
	Status        string    `bun:"status_,notnull" json:"status"`
	CreatedAt     time.Time `bun:"created_at,default:current_timestamp" json:"created_at"`
	AssignedBy    *int64    `bun:"assigned_by,nullzero" json:"assigned_by,omitempty"`
}

type AppointmentRequest struct {
	PatientID int64  `bun:"user_id" json:"patient_id"`
	Status    string `bun:"status" json:"status_"`
}

type AssignRequest struct {
	StaffID *int64 `bun:"user_id,nullzero" json:"user_id,omitempty"`
	// StaffRole *int64 `bun:"user_role,nullzero" json:"user_role,omitempty"`
	Status string `bun:"status_" json:"status_"`
}

type LoginRequest struct {
	Email    string `bun:"email" json:"email"`
	Password string `bun:"password" json:"pass_word"`
}

type Claims struct {
	ID    int64  `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.StandardClaims
}
