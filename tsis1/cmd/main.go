package main

import (
	"tsis1/pkg"
	"net/http"
	"github.com/gorilla/mux"
)


func main() {
	router := mux.NewRouter()

	router.HandleFunc("/health-check", pkg.HealthCheck).Methods("GET")
	router.HandleFunc("/characters", pkg.Characters).Methods("GET")
	router.HandleFunc("/characters/{id:[0-9]+}", pkg.CharacterByID).Methods("GET")

	http.Handle("/", router)
	http.ListenAndServe(":8080", router)
}