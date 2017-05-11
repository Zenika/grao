// The controller subpackage exposes endpoint handlers
// used to call search service through REST API
package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Zenika/RAO/log"
	"github.com/Zenika/RAO/search"
	"github.com/gorilla/mux"
)

func SearchHandler(searchService *search.SearchService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		handleSearch(w, r, searchService)
	}
}

func SettingsHandler(searchService *search.SearchService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		handleConfig(w, r, searchService)
	}
}

func handleConfig(w http.ResponseWriter, r *http.Request, searchService *search.SearchService) {
	vars := mux.Vars(r)
	index := vars["index"]
	var settings search.Settings
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&settings)
	log.Error(err, log.ERROR)
	searchService.Configure(index, settings)
}

func handleSearch(w http.ResponseWriter, r *http.Request, searchService *search.SearchService) {
	vars := mux.Vars(r)
	index := vars["index"]
	var query search.Query
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&query)
	if nil != err {
		log.Error(err, log.ERROR)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("There was an error: %v", err)))
		return
	}
	queryRes, err := searchService.Search(index, query)
	response, err := json.Marshal(queryRes.Data)
	if err == nil {
		w.Write([]byte(response))
	} else {
		log.Error(err, log.ERROR)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("There was an error: %v", err)))
	}
}
