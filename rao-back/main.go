package main

import (
	"github.com/Zenika/RAO/algolia"
	"github.com/Zenika/RAO/dropbox"
	"github.com/Zenika/RAO/docd"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"net/http"
	"io/ioutil"
	"log"
	"io"
	"os"
)

var logFile string = os.Getenv("RAO_LOG_FILE")
var documents []dropbox.DbxDocument

func etl(w http.ResponseWriter, r *http.Request) {
	root := os.Getenv("RAO_DBX_ROOT")
	dropbox.Walk(root, func(res io.ReadCloser, doc dropbox.DbxDocument){
		buffer, err := ioutil.ReadAll(res)
		defer res.Close()
		check(err)
		content, _, err := docd.Convert(buffer, doc.Mime)
		check(err)
		doc.Content = string(content[:])
		log.Println(doc.Content)
		documents = append(documents, doc)
	});
	algolia.Push(documents)
}

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
		//AllowedOrigins:   []string{"http://localhost:*"},
		AllowedMethods:   []string{"GET", "POST", "DELETE", "OPTION", "PUT"},
		AllowCredentials: true,
	})

	r := mux.NewRouter()
	r.HandleFunc("/api/v1", etl)

	handler := c.Handler(r)

	http.ListenAndServe(":8090", handler)

}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
