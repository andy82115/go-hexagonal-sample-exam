package domain

import (
	"time"
)

type UserRole string

const (
	Premium   UserRole = "premium"
	Normal UserRole = "normal"
)

type User struct {
	ID        uint64
	Name      string
	Email     string
	Password  string
	Role      UserRole
	CreatedAt time.Time
	UpdatedAt time.Time
}