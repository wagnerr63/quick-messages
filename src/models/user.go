package models

import (
	"database/sql"
	"fmt"
	"math/rand"
	"time"
)

// User struct
type User struct {
	ID        int        `json:"id"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

// AllUsers mock
var AllUsers = []User{
	User{ID: rand.Intn(100), Name: "Wagner", Email: "wagner@mail.com", Password: "teste123"},
}

// Create add a new User
func Create(user User) {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/quick_messages")
	if err != nil {
		panic(err.Error())
	}
	newUser, err := db.Prepare("INSERT INTO users(name, email, password) VALUES (?,?,?);")
	if err != nil {
		panic(err.Error())
	}

	userID, err := newUser.Exec(user.Name, user.Email, user.Password)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	fmt.Println(userID)
}
