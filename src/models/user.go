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
	Image     string     `json:"image"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	Level     int        `json:"level"`  // 0 - Normal | 1 - MASTER
	Status    int        `json:"status"` // 0 - Inative | 1 - Ative
}

// FindAll returns all users
func FindAll() []User {
	db := db.ConnectDatabase()

	selectAllProducts, err := db.Query("SELECT * FROM users ORDER BY created_at ASC;")
	if err != nil {
		panic(err.Error())
	}

	u := User{}
	users := []User{}

	for selectAllProducts.Next() {
		var id, status, level int
		var name, email, password, image string
		var createdAt, updatedAt *time.Time

		err = selectAllProducts.Scan(&id, &name, &email, &password, &image, &createdAt, &updatedAt, &level, &status)
		if err != nil {
			panic(err.Error())
		}

		u.ID = id
		u.Name = name
		u.Email = email
		u.Image = image
		u.CreatedAt = createdAt
		u.UpdatedAt = updatedAt
		u.Level = level
		u.Status = status

		users = append(users, u)
	}

	defer db.Close()
	return users
}

// Create add a new User
func Create(user User) {
	db := db.ConnectDatabase()

	insertUser, err := db.Prepare("INSERT INTO users(name, email, password, image, level) VALUES (?,?,?,?,?);")
	if err != nil {
		panic(err.Error())
	}

	userID, err := insertUser.Exec(user.Name, user.Email, user.Password, user.Image, user.Level)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	fmt.Println(userID)
}
