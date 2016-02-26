package handlers

import (
	"encoding/json"
	"github.com/odacremolbap/k8stest/model"
	"log"
	"net/http"
)

// AppInfoHandler takes care of container info requests
func AppInfoHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received request for %s", r.URL.Path)

	appInfo := model.GetAppInfo()
	appJSON, err := json.Marshal(appInfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(appJSON)
}
