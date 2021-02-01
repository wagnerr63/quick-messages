package models

import (
	"time"
)

// User struct
type User struct {
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

// AllUsers mock
var AllUsers = []User{
	User{Name: "Wagner", Email: "wagner@mail.com", Password: "teste123"},
	User{Name: "Rafael", Email: "rafael@mail.com", Password: "teste123"},
}
