package main

import (
	"github.com/Zenika/RAO/dropbox"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"net/http"
	"log"
	"os"
)

var logFile string = os.Getenv("RAO_LOG_FILE");

/**
 * TODO :
 * - P1: Lire le chemin racine des documents dans le fichier de config (ou variable d'env)
 * - P1: Au lancement : parcourir tous les documents et les indexer dans Algolia
 * - P1: Stocker en BDD les meta des fichiers déjà indexés (id, chemin relatif, nom)
 * - P1: Tous les soirs, parcourir les nouveaux documents et les indexer dans Algolia + maj la BDD
 *
 * - P2: services pour rechercher dans Algolia
 * - P2: offrir une IHM pour rechercher les documents via mots clés
 */

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
		AllowedOrigins:   []string{"*"},
		//AllowedOrigins:   []string{"http://localhost:*"},
		AllowedMethods:   []string{"GET", "POST", "DELETE", "OPTION", "PUT"},
		AllowCredentials: true,
	})

	r := mux.NewRouter()
	r.HandleFunc("/api/v1", dropbox.GetRootFolder)

	handler := c.Handler(r)

	http.ListenAndServe(":8090", handler)

}
