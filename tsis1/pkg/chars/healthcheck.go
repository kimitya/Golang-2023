package characters

import (
	"fmt"
	"net/http"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hi! My name is Anita c: My favorite movie is Ratatoulie, so it's my API of characters of the movie <з")
}