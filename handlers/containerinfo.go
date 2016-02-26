package handlers

import (
	"encoding/json"
	"github.com/odacremolbap/k8stest/model"
	"log"
	"net/http"
)

// ContainerInfoHandler takes care of container info requests
func ContainerInfoHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received request for %s", r.URL.Path)

	containerInfo, err := model.GetContainerInfo()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	appJSON, err := json.Marshal(containerInfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(appJSON)
}
