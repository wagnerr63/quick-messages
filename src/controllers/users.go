package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"../models"
)

// UsersIndex list all users
func UsersIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: /users")
	users := models.FindAll()

	json.NewEncoder(w).Encode(users)
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

	user.Password = GetMD5Hash(user.Password)

	models.Create(user)
	json.NewEncoder(w).Encode(user)
}

// GetMD5Hash cripto
func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
