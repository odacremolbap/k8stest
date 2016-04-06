package handlers

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// ErrOutHandler exits the app with output code 2
func ErrOutHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received request for %s", r.URL.Path)

	defer os.Exit(1)
	fmt.Fprintf(w, "Exiting server!")
}
