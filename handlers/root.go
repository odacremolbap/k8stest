package handlers

import (
	"log"
	"net/http"

	"github.com/odacremolbap/k8stest/model"
)

// RootHandler takes care of root site requests
func RootHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received request for %s", r.URL.Path)

	t, err := model.GetIndexPage()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	containerInfo, err := model.GetContainerInfo()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	indexPage := model.IndexPage{
		AppInfo:       model.GetAppInfo(),
		ContainerInfo: *containerInfo,
	}

	log.Print("[Root]Executing template")
	err = t.ExecuteTemplate(w, "index", indexPage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
