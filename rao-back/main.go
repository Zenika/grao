package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Zenika/RAO/conv"
	"github.com/Zenika/RAO/conv/docd"
	bdcService "github.com/Zenika/RAO/document/bdc/service"
	raoService "github.com/Zenika/RAO/document/rao/service"
	"github.com/Zenika/RAO/log"
	"github.com/Zenika/RAO/auth/auth0"
	"github.com/Zenika/RAO/search"
	"github.com/Zenika/RAO/search/algolia"
	searchController "github.com/Zenika/RAO/search/controller"
	"github.com/Zenika/RAO/tree"
	"github.com/Zenika/RAO/tree/dropbox"
	"github.com/gorilla/mux"
	"github.com/robfig/cron"
	"github.com/rs/cors"
)

/* INIT SERVICES IMPLEMENTATIONS */
var treeService = tree.New(dropbox.New())
var convService = conv.New(docd.New())
var searchService = search.New(algolia.New())

func main() {
	/* INIT LOGGING */
	log.Init()
	defer log.Close()
	log.Info("Application started")
	/* INIT SCHEDULLER */
	cronExp := os.Getenv("GRAO_POLL_EVERY")
	if len(cronExp) == 0 {
		cronExp = "@daily" // equivalent to 0 0 0 * * *
	}
	cron := cron.New()
	cron.AddFunc(cronExp, func() {
		root := fmt.Sprintf("/%v", os.Getenv("GRAO_DBX_ROOT"))
		bdcService := bdcService.New(*searchService, *treeService)
		raoService := raoService.New(*searchService, *convService, *treeService)
		pairs := [][]interface{}{{bdcService.DocFilter, bdcService.DocHandler}, {raoService.DocFilter, raoService.DocHandler}}
		treeService.Poll(root, pairs)
	})
	cron.Start()
	/* INIT HTTP CONTROLLER */
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "DELETE", "OPTION", "PUT"},
		AllowedHeaders: []string{"authorization", "content-type"},
		AllowCredentials: true,
	})
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/{index}/search", searchController.SearchHandler(searchService)).
		Methods("POST")
	r.HandleFunc("/api/v1/{index}/settings", searchController.SettingsHandler(searchService)).
		Methods("POST")
	test := os.Getenv("AUTH0_JWKS_URI")
	fmt.Println(test)
	fmt.Println(os.Getenv("AUTH0_ISSUER"))
	fmt.Println(os.Getenv("AUTH0_AUDIENCE"))
	auth := auth0.New(
		os.Getenv("AUTH0_JWKS_URI"),
		os.Getenv("AUTH0_ISSUER"),
		os.Getenv("AUTH0_AUDIENCE"),
	)
	auth0Middleware := auth.UserAuthenticatedMiddleware
	handler := auth0Middleware(c.Handler(r))
	http.ListenAndServe(":8090", handler)
}
