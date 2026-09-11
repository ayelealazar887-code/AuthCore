package models

import (
	"time"

	"github.com/google/uuid"
)

type UserType string

const (
	Admin    UserType = "admin"
	Employee UserType = "employee"
)

type User struct {
	ID           uuid.UUID  `json:"id" db:"id"`
	FirstName    string     `json:"first_name" db:"first_name"`
	LastName     string     `json:"last_name" db:"last_name"`
	Password     string     `json:"password" db:"password"`
	Email        string     `json:"email" db:"email"`
	Phone        string     `json:"phone" db:"phone"`
	Token        string     `json:"token" db:"token"`
	UserType     UserType   `json:"user_type" db:"user_type"`
	RefreshToken string     `json:"refresh_token" db:"refresh_token"`
	UserID       uuid.UUID  `json:"user_id" db:"user_id"`
	CreatedAt    time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at" db:"updated_at"`
}