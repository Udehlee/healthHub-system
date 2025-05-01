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
}

type Conn struct {
	DB  *bun.DB
	Ctx context.Context
}

func NewConn(db *bun.DB) Conn {
	ctx := context.Background()
	return Conn{
		DB:  db,
		Ctx: ctx,
	}
}

// Save saves new user details to DB
func (c *Conn) Save(user *models.User) error {
	_, err := c.DB.NewInsert().
		Model(user).Returning("user_Id, email, role").
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
		Where("email = ?", email).Limit(1).
		Scan(c.Ctx)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("error fetching user: %w", err)
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
		Model(appointment).Returning("patient_Id, status_").
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
		Set("staff_role = ?", appointment.StaffRole).
		Set("assigned_by = ?", appointment.AssignedBy).
		Where("appointment_id = ?", appointmentID).
		Exec(c.Ctx)

	if err != nil {
		return fmt.Errorf("failed to assign appoinment details: %w", err)
	}
	return nil
}
