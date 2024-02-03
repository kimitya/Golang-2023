package main

import (
	"tsis1/pkg/chars"
	"net/http"
	"github.com/gorilla/mux"
)


func main() {
	router := mux.NewRouter()

	router.HandleFunc("/health-check", characters.HealthCheck).Methods("GET")
	router.HandleFunc("/characters", characters.Characters).Methods("GET")
	router.HandleFunc("/characters/{id:[0-9]+}", characters.CharacterByID).Methods("GET")

	http.Handle("/", router)
	http.ListenAndServe(":8080", router)
}