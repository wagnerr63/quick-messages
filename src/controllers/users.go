package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../models"
)

// UsersIndex list all users
func UsersIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: /users")
	json.NewEncoder(w).Encode(models.AllUsers)
}

// UsersStore list all users
func UsersStore(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: /users Method: POST")

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	if user.Name == "" {
		http.Error(w, "Name is required", 400)
		return
	}

	if user.Email == "" {
		http.Error(w, "E-mail is required", 400)
		return
	}

	if user.Password == "" {
		http.Error(w, "Password is required", 400)
		return
	}

	models.AllUsers = append(models.AllUsers, user)
	json.NewEncoder(w).Encode(user)
}
