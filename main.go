package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"

	"github.com/odacremolbap/k8stest/handlers"
)

var hits int
var hi string
var appInfo string
var mtx sync.Mutex

const defaultPort int = 9876

func main() {
	log.Print("Initializing server")

	createRoutes()
	port := getListeningPort()

	// initialize info
	var err error

	if err != nil {
		log.Fatalf("Error retrieving host info: %s", err.Error())
	}

	log.Printf("Listening on port %d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

// gets default port from environment, and falls back to default
func getListeningPort() int {

	if argPort := os.Getenv("K8STEST_PORT"); argPort != "" {
		if i, err := strconv.Atoi(argPort); err == nil {
			if i != 0 {
				log.Printf("Getting port from K8STEST_PORT = %s", argPort)
				return i
			}
		} else {
			log.Printf("Environment variable K8STEST_PORT = %s isn't a valid number", argPort)
		}
	}

	log.Printf("Using default port %d", defaultPort)
	return defaultPort
}

// create routes
func createRoutes() {
	http.HandleFunc("/", handlers.RootHandler)
	http.HandleFunc("/containerinfo", handlers.ContainerInfoHandler)
	http.HandleFunc("/appinfo", handlers.AppInfoHandler)
	http.HandleFunc("/callenv", handlers.CallEnvServiceHandler)
	http.HandleFunc("/exit0", handlers.CleanOutHandler)
	http.HandleFunc("/exit1", handlers.ErrOutHandler)

}
