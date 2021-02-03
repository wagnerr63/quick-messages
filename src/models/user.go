package models

import (
	"fmt"
	"time"

	"../db"
)

// User struct
type User struct {
	ID        int        `json:"id"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	Status    int        `json:"status"`
}

// FindAll returns all users
func FindAll() []User {
	db := db.ConnectDatabase()

	selectAllProducts, err := db.Query("SELECT * FROM users ORDER BY id ASC;")
	if err != nil {
		panic(err.Error())
	}

	u := User{}
	users := []User{}

	for selectAllProducts.Next() {
		var id, status int
		var name, email, password string
		var createdAt, updatedAt *time.Time

		err = selectAllProducts.Scan(&id, &name, &email, &password, &createdAt, &updatedAt, &status)
		if err != nil {
			panic(err.Error())
		}

		u.ID = id
		u.Name = name
		u.Email = email
		u.CreatedAt = createdAt
		u.UpdatedAt = updatedAt
		u.Status = status

		users = append(users, u)
	}

	defer db.Close()
	return users
}

// Create add a new User
func Create(user User) {
	db := db.ConnectDatabase()

	insertUser, err := db.Prepare("INSERT INTO users(name, email, password) VALUES (?,?,?);")
	if err != nil {
		panic(err.Error())
	}

	userID, err := insertUser.Exec(user.Name, user.Email, user.Password)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	fmt.Println(userID)
}
