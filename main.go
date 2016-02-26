package main

import (
	"fmt"
	"github.com/odacremolbap/k8stest/handlers"
	//	"github.com/odacremolbap/k8stest/model"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"sync"
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
	hi, err = getHostInfo()

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
}

func getHostInfo() (string, error) {
	log.Print("getHostInfo")
	interfaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}

	hostName, err := os.Hostname()
	if err != nil {
		return "", err
	}

	hostInfo := fmt.Sprintf("hostname: %s", hostName)

	for _, iface := range interfaces {
		addrs, err := iface.Addrs()
		if err != nil {
			return "", err
		}
		for i := range addrs {
			hostInfo = fmt.Sprintf("%s\n%s-address-%d: %s", hostInfo, iface.Name, i, addrs[i].String())
		}
	}

	return hostInfo, nil
}

// func getAppInfo() string {
// 	return fmt.Sprintf("k8stest-version: %s\nk8stest-date: %s\nk8stest-commit: %s", appinfo.Version, appinfo.Date, appinfo.GitCommit)
// }

// func hello(w http.ResponseWriter, r *http.Request) {
// 	log.Printf("Received request for %s", r.URL.Path)
// 	mtx.Lock()
// 	hits++
// 	mtx.Unlock()
// 	fmt.Fprintf(w, "%s\nhits: %d\n\n%s", hi, hits, appInfo)
// }
