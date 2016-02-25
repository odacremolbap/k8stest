package main

import (
	"fmt"
	"github.com/odacremolbap/k8stest/appinfo"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"sync"
)

var Date string
var hits int
var hi string
var appInfo string
var mtx sync.Mutex

func main() {
	log.Print("Initializing server")
	var err error
	hi, err = getHostInfo()
	appInfo = getAppInfo()

	if err != nil {
		log.Fatalf("Error retrieving host info: %s", err.Error())
	}

	http.HandleFunc("/", hello)

	port := 9876
	if argPort := os.Getenv("K8STEST_PORT"); argPort != "" {
		if i, err := strconv.Atoi(argPort); err == nil {
			if i != 0 {
				port = i
			}
		}
	}

	log.Printf("Listening on port %d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
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

func getAppInfo() string {
	return fmt.Sprintf("k8stest-version: %s\nk8stest-date: %s\nk8stest-commit: %s", appinfo.Version, appinfo.Date, appinfo.GitCommit)
}

func hello(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received request for %s", r.URL.Path)
	mtx.Lock()
	hits++
	mtx.Unlock()
	fmt.Fprintf(w, "%s\nhits: %d\n\n%s", hi, hits, appInfo)
}
