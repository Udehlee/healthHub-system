package models

import (
	"time"

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

type Appointment struct {
	AppointmentID int64     `bun:"appointment_id,pk,autoincrement"`
	PatientID     int64     `bun:"patient_id"`
	StaffID       *int64    `bun:"staff_id,nullzero"`
	StaffRole     *int64    `bun:"staff_role,nullzero"`
	Status        string    `bun:"status_"`
	CreatedAt     time.Time `bun:"created_at,default:current_timestamp"`
	AssignedBy    *int64    `bun:"assigned_by,nullzero"`
}

type AppointmentRequest struct {
	PatientID int64  `bun:"patient_id" json:"patient_id"`
	Status    string `bun:"status" json:"status"`
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
