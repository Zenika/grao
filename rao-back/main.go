package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Zenika/RAO/controller"
	"github.com/gorilla/mux"
	"github.com/robfig/cron"
	"github.com/rs/cors"
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

	cronExp := os.Getenv("RAO_POLL_EVERY")
	if len(cronExp) == 0 {
		cronExp = "@daily" // equivalent to 0 0 0 * * *
	}
	cron := cron.New()
	cron.AddFunc(cronExp, controller.Poll)
	cron.Start()
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "DELETE", "OPTION", "PUT"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"Authorization"},
	})
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/search", controller.Search)
	handler := c.Handler(r)
	http.ListenAndServe(":8090", handler)
}
