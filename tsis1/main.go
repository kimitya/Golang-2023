package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
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
	fmt.Fprintf(w, "Hi! My name is Anita c: My favorite movie is Ratatoulie, so it's my API of characters ofthe movie <з")
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

func Characters(w http.ResponseWriter, r *http.Request) {
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

func CharacterByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	characterIDStr, ok := vars["id"]
	if !ok {
		http.Error(w, "Missing character ID", http.StatusBadRequest)
		return
	}

	characterID, err := strconv.Atoi(characterIDStr)
	if err != nil {
		http.Error(w, "Invalid character ID", http.StatusBadRequest)
		return
	}

	allCharacters := prepareResponse()

	var foundCharacter Character
	for _, character := range allCharacters {
		if character.ID == characterID {
			foundCharacter = character
			break
		}
	}

	if foundCharacter.ID == 0 {
		http.Error(w, "Character not found", http.StatusNotFound)
		return
	}

	jsonResponse, err := json.Marshal(foundCharacter)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)

	w.Write(jsonResponse)
}



func main() {
	router := mux.NewRouter()

	router.HandleFunc("/health-check", HealthCheck).Methods("GET")
	router.HandleFunc("/characters", Characters).Methods("GET")
	router.HandleFunc("/characters/{id:[0-9]+}", CharacterByID).Methods("GET")

	http.Handle("/", router)
	http.ListenAndServe(":8080", router)
}