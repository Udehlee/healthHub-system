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
	Save(user models.User) error
	CheckEmail(email string) (*models.User, error)
}

type Conn struct {
	DB *bun.DB
}

func NewConn(db *bun.DB) Conn {
	return Conn{
		DB: db,
	}
}

// Save saves new user details to DB
func (c *Conn) Save(user *models.User) error {
	ctx := context.Background()

	_, err := c.DB.NewInsert().
		Model(user).Returning("user_Id, email, role").
		Exec(ctx)

	if err != nil {
		return fmt.Errorf("failed to insert user: %w", err)
	}

	return nil
}

// CheckEmail checks for already used Email
func (c *Conn) CheckEmail(email string) (*models.User, error) {
	ctx := context.Background()
	user := new(models.User)

	err := c.DB.NewSelect().
		Model(user).
		Where("email = ?", email).Limit(1).
		Scan(ctx)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("error fetching user: %w", err)
	}

	return user, nil
}
