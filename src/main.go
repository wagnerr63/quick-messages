package main

import (
	"fmt"
	"log"
	"net/http"

	"./controllers"

	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "REST API with GO")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	// creates a new instance of a mux router
	routes := mux.NewRouter().StrictSlash(true)

	routes.HandleFunc("/", homePage)
	routes.HandleFunc("/users", controllers.UsersIndex).Methods("GET")
	routes.HandleFunc("/users", controllers.UsersStore).Methods("POST")

	log.Fatal(http.ListenAndServe(":10000", routes))
}

func main() {
	fmt.Println("🔥 Server started at http://localhost:10000")
	handleRequests()
}
