package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

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

// Show All Sushi
func getRolls(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rolls)
}

// Show Single Sushi
func getRoll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range rolls {
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

// add Single Sushi
func createRoll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newRoll  Roll
	json.NewDecoder(r.Body).Decode(&newRoll)
	newRoll.ID = strconv.Itoa(len(rolls) + 1)
	rolls  = append(rolls, newRoll)
	json.NewEncoder(w).Encode(newRoll)
}

// update Single Sushi
func updateRoll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for i, item := range rolls{
		if item.ID == params["id"] {
			rolls = append(rolls[:i], rolls[i+1:]...)
			var newRoll Roll
			json.NewDecoder(r.Body).Decode(&newRoll)
			newRoll.ID = params["id"]
			rolls  = append(rolls, newRoll)
			json.NewEncoder(w).Encode(newRoll)
			return
		}
	}
}

// delete Single Sushi
func deleteRoll(w http.ResponseWriter, r *http.Request) {

}

func main() {
	// Generate Mock Data
	rolls = append(rolls,
		Roll{
			ID: "1",
			Name: "Salmon Roll",
			Description: "crab stick, tamago sushi with salmon and cheese",
			Ingredient: "Salmon, Nori, Soy Sauce, Rice",
		},
		Roll{
			ID: "2",
			Name: "California Roll",
			Description: "crab stick, tamago sushi with california and cheese",
			Ingredient: "California, Nori, Soy Sauce, Rice",
		},
	)


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
