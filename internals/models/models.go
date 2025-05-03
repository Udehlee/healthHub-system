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
	RoleID    int64  `bun:"role_id,notnull" json:"role_id"`
	Gender    string `bun:"gender" json:"gender"`
	Address   string `bun:"user_address" json:"user_address"`

	Role *Role `bun:"rel:belongs-to,join:role_id=role_id" json:"role,omitempty"`
}

type Appointment struct {
	bun.BaseModel `bun:"table:appointments"`

	AppointmentID int64     `bun:"appointment_id,pk,autoincrement" json:"appointment_id"`
	PatientID     int64     `bun:"patient_id,notnull" json:"patient_id"`
	StaffID       *int64    `bun:"staff_id,nullzero" json:"staff_id,omitempty"`
	StaffRole     *int64    `bun:"staff_role,nullzero" json:"staff_role,omitempty"`
	Status        string    `bun:"status_,notnull" json:"status"`
	CreatedAt     time.Time `bun:"created_at,default:current_timestamp" json:"created_at"`
	AssignedBy    *int64    `bun:"assigned_by,nullzero" json:"assigned_by,omitempty"`
}

type Role struct {
	bun.BaseModel `bun:"table:roles"`

	RoleID   int64  `bun:"role_id,pk,autoincrement" json:"role_id"`
	RoleName string `bun:"role_name,unique,notnull" json:"role_name"`
}

type AppointmentRequest struct {
	PatientID int64  `bun:"user_id" json:"patient_id"`
	Status    string `bun:"status" json:"status_"`
}

type AssignRequest struct {
	StaffID   *int64 `bun:"user_id,nullzero" json:"user_id,omitempty"`
	StaffRole *int64 `bun:"user_role,nullzero" json:"user_role,omitempty"`
	Status    string `bun:"status_" json:"status_"`
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
