package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Zenika/RAO/log"
	"github.com/Zenika/RAO/search"
	"github.com/Zenika/RAO/search/algolia"
	"github.com/gorilla/mux"
)

func SearchHandler(searchService *search.SearchService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		handleSearch(w, r, searchService)
	}
}

func SettingsHandler(searchService *search.SearchService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		handlerConfig(w, r, searchService)
	}
}

func handlerConfig(w http.ResponseWriter, r *http.Request, searchService *search.SearchService) {

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
	queryRes, err := searchService.Search(query, algolia.SearchOptions{Index: index})
	response, err := json.Marshal(queryRes.Data)
	if err == nil {
		w.Write([]byte(response))
	} else {
		log.Error(err, log.ERROR)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("There was an error: %v", err)))
	}
}
