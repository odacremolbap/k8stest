package handlers

import (
	"fmt"
	"log"
	"net/http"
)

// RootHandler takes care of root site requests
func RootHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received request for %s", r.URL.Path)
	fmt.Fprint(w, "TODO page. Please check /containerinfo and /appinfo")
}
