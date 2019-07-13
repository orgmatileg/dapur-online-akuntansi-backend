package model

import (
	"time"
)

// User Struct
type User struct {
	UserID       int64     `json:"user_id"`
	UserRoleID   int64     `json:"user_role_id"`
	UserRoleName string    `json:"user_role_name"`
	Email        string    `json:"email"`
	Password     string    `json:"password"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	PhotoProfile string    `json:"photo_profile"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// Users / User list
type Users []User

// NewUser func
func NewUser() *User {
	return &User{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
