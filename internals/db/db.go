package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/Udehlee/healthHub-System/internals/models"
	"github.com/uptrace/bun"
)

type Store interface {
	Save(user *models.User) error
	CheckEmail(email string) (*models.User, error)
	GetAllUsers() ([]*models.User, error)
	SaveAppointment(appointment *models.Appointment) error
	AssignStaff(appointmentID int64, appointment *models.Appointment) error
	GetAssignedAppointments() ([]*models.Appointment, error)
}

type Conn struct {
	DB  *bun.DB
	Ctx context.Context
}

func NewConn(db *bun.DB) *Conn {
	ctx := context.Background()
	return &Conn{
		DB:  db,
		Ctx: ctx,
	}
}

// Save saves new user details to DB
func (c *Conn) Save(user *models.User) error {
	_, err := c.DB.NewInsert().
		Model(user).Returning("user_Id, email, user_role").
		Exec(c.Ctx)

	if err != nil {
		return fmt.Errorf("failed to insert user: %w", err)
	}

	return nil
}

// CheckEmail checks for already used Email
func (c *Conn) CheckEmail(email string) (*models.User, error) {
	user := new(models.User)

	err := c.DB.NewSelect().
		Model(user).
		Column("user_id", "email", "pass_word", "user_role").
		Where("email = ?", email).
		Limit(1).
		Scan(c.Ctx)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // email not taken
		}
		return nil, fmt.Errorf("error checking email: %w", err)
	}

	return user, nil
}

// GetAllUsers retieves all users from database
func (c *Conn) GetAllUsers() ([]*models.User, error) {
	var users []*models.User

	err := c.DB.NewSelect().
		Model(&users).
		Scan(c.Ctx)
	if err != nil {
		return []*models.User{}, fmt.Errorf("error fetching all users: %w", err)
	}
	return users, nil

}

// SaveAppointment Saves appointment details
func (c *Conn) SaveAppointment(appointment *models.Appointment) error {
	_, err := c.DB.NewInsert().
		Model(appointment).Returning("patient_id, status_").
		Exec(c.Ctx)

	if err != nil {
		return fmt.Errorf("failed to insert appoinment details: %w", err)
	}

	return nil
}

// AssignStaff updates the appointment table  with the given staff ID,Role and assignedby fields
func (c *Conn) AssignStaff(appointmentID int64, appointment *models.Appointment) error {
	_, err := c.DB.NewUpdate().
		Model(appointment).
		Set("staff_id = ?", appointment.StaffID).
		// Set("staff_role = ?", appointment.StaffRole).
		Set("assigned_by = ?", appointment.AssignedBy).
		Where("appointment_id = ?", appointmentID).
		Exec(c.Ctx)

	if err != nil {
		return fmt.Errorf("failed to assign appointment details: %w", err)
	}
	return nil
}

// GetAssignedAppointments returns all assigned appointment details from db
func (c *Conn) GetAssignedAppointments() ([]*models.Appointment, error) {
	var appointments []*models.Appointment

	err := c.DB.NewSelect().
		Model(&appointments).
		Where("status = ?", "assigned").
		Scan(c.Ctx)

	if err != nil {
		return []*models.Appointment{}, fmt.Errorf("failed to get assigned appoinment: %w", err)
	}
	return appointments, nil
}
