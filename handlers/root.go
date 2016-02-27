package handlers

import (
	"github.com/odacremolbap/k8stest/model"
	"log"
	"net/http"
)

// RootHandler takes care of root site requests
func RootHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("[Root]Received request for %s", r.URL.Path)

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
