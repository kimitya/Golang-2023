package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Response struct {
	Characters []Character `json:"persons"`
}

type Character struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Role      string `json:"role"`
	IsRat     bool   `json:"is_rat"`
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	//specify status code
	w.WriteHeader(http.StatusOK)
  
  //update response writer 
	fmt.Fprintf(w, "API is up and running")
}

func prepareResponse() []Character {
	var characters []Character

	characters = append(characters, Character{ID: 1, Name: "Remy", Role: "Main Chef", IsRat: true})
	characters = append(characters, Character{ID: 2, Name: "Linguini", Role: "Chef's Assistant", IsRat: false})
	characters = append(characters, Character{ID: 3, Name: "Colette", Role: "Chef", IsRat: false})
	characters = append(characters, Character{ID: 4, Name: "Anton Ego", Role: "Food Critic", IsRat: false})
	characters = append(characters, Character{ID: 5, Name: "Skinner", Role: "Head Chef", IsRat: false})
	characters = append(characters, Character{ID: 6, Name: "Emile", Role: "Remy's Brother", IsRat: true})
	characters = append(characters, Character{ID: 7, Name: "Gusteau", Role: "Former Chef", IsRat: false})
	characters = append(characters, Character{ID: 8, Name: "Collette Tatou", Role: "Chef", IsRat: false})
	characters = append(characters, Character{ID: 9, Name: "Horst", Role: "Chef", IsRat: false})
	characters = append(characters, Character{ID: 10, Name: "Django", Role: "Remy's Father", IsRat: true})
  
	return characters
}

func Persons(w http.ResponseWriter, r *http.Request) {
	//declare response variable
	var response Response
  
	//Retrieve person details
	characters := prepareResponse()
  
	//assign person details to response
	response.Characters = characters
  
	//update content type
	w.Header().Set("Content-Type", "application/json")
  
	//specify HTTP status code
	w.WriteHeader(http.StatusOK)
  
	//convert struct to JSON
	jsonResponse, err := json.Marshal(response)
	if err != nil {
	 return
	}
  
	//update response
	w.Write(jsonResponse)
}

func main() {
	// Create a new router
	router := mux.NewRouter()

	// Specify endpoints, handler functions, and HTTP methods
	router.HandleFunc("/health-check", HealthCheck).Methods("GET")
	router.HandleFunc("/characters", Persons).Methods("GET")

	// Start and listen to requests
	http.Handle("/", router)
	http.ListenAndServe(":8080", router)
}