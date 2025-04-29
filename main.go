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

func main() {
	// Init Router
	router := mux.NewRouter()

	// Handle Endpoint/Routing
	router.HandleFunc("/sushi").Methods("GET")
	router.HandleFunc("/sushi/{id}").Methods("GET")
	router.HandleFunc("/sushi").Methods("POST")
	router.HandleFunc("/sushi/{id}").Methods("POST")
	router.HandleFunc("/sushi/{id}").Methods("DELETE")

	log.Fatal(http.ListenAndServe(":5000", router))
}
