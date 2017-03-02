package main

import (
	"github.com/Zenika/RAO/rao"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"net/http"
	"log"
	"os"
)

var logFile string = os.Getenv("RAO_LOG_FILE")

func main() {

	if len(logFile) == 0 {
		logFile = "rao.log"
	}
	f, err := os.OpenFile(logFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	log.SetOutput(f)
	log.Println("Application started")

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "DELETE", "OPTION", "PUT"},
		AllowCredentials: true,
	})

	r := mux.NewRouter()
	r.HandleFunc("/api/v1/index", rao.IndexAllDropBoxDocuments)

	handler := c.Handler(r)

	http.ListenAndServe(":8090", handler)

}