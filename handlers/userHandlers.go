package handlers

import (
	"fmt"
	"net/http"
)

// Index : list users
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome!")
}

func Create(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Creating")
}
