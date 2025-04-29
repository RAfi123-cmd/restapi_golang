package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Roll is a Model for sushi API
type Roll struct {
	ID          string `json: "id"`
	Name        string `json: "name"`
	Description string `json: "description"`
	Ingredient  string `json: ingredient`
}
// Init Rolls var as a slice
var rolls []Roll

// Index 
func getRolls(w http.ResponseWriter, r *http.Request)  {
	
}
// Show Single Sushi 
func getRoll(w http.ResponseWriter, r *http.Request)  {
	
}
// add Single Sushi 
func createRoll(w http.ResponseWriter, r *http.Request)  {
	
}

// update Single Sushi 
func updateRoll(w http.ResponseWriter, r *http.Request)  {
	
}
// delete Single Sushi 
func deleteRoll(w http.ResponseWriter, r *http.Request)  {
	
}

func main() {
	// Init Router
	router := mux.NewRouter()

	// Handle Endpoint/Routing
	router.HandleFunc("/sushi", getRolls).Methods("GET")
	router.HandleFunc("/sushi/{id}", getRoll).Methods("GET")
	router.HandleFunc("/sushi", createRoll).Methods("POST")
	router.HandleFunc("/sushi/{id}", updateRoll).Methods("POST")
	router.HandleFunc("/sushi/{id}", deleteRoll).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":5000", router))
}
