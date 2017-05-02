package main

import (
	"github.com/Zenika/RAO/controller"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
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

	controller.Poll(nil, nil)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "DELETE", "OPTION", "PUT"},
		AllowCredentials: true,
	})

	r := mux.NewRouter()
	r.HandleFunc("/api/v1/walk", controller.Walk)
	r.HandleFunc("/api/v1/search", controller.Search)
	// r.HandleFunc("/api/v1/poll", controller.Poll)

	handler := c.Handler(r)
	http.ListenAndServe(":8090", handler)

}
