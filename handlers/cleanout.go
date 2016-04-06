package handlers

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// CleanOutHandler exits the app with output code 0
func CleanOutHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received request for %s", r.URL.Path)

	defer os.Exit(0)
	fmt.Fprintf(w, "Exiting server!")
}
