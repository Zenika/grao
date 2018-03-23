package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Zenika/rao/rao-back/conv"
	"github.com/Zenika/rao/rao-back/conv/docd"
	bdcService "github.com/Zenika/rao/rao-back/document/bdc/service"
	raoService "github.com/Zenika/rao/rao-back/document/rao/service"
	"github.com/Zenika/rao/rao-back/log"
	"github.com/Zenika/rao/rao-back/auth/auth0"
	"github.com/Zenika/rao/rao-back/search"
	"github.com/Zenika/rao/rao-back/search/algolia"
	searchController "github.com/Zenika/rao/rao-back/search/controller"
	"github.com/Zenika/rao/rao-back/tree"
	"github.com/Zenika/rao/rao-back/tree/dropbox"
	"github.com/gorilla/mux"
	"github.com/robfig/cron"
	"github.com/rs/cors"
)


/* INIT SERVICES IMPLEMENTATIONS */
var treeService = tree.New(dropbox.New())
var convService = conv.New(docd.New())
var searchService = search.New(algolia.New())

/*
GRAO_APP_PORT=8090
GRAO_DBX_KEY=dropbox_key
GRAO_DBX_SECRET=dropbox_secret
GRAO_DBX_TOKEN=secret_token
GRAO_DBX_ROOT=/
GRAO_DBX_CURSOR=/cursor
GRAO_ALGOLIA_ID=algolia_api_client_id
GRAO_ALGOLIA_KEY=algolia_api_key
GRAO_LOG_FILE=rao.log
GRAO_LOG_LEVEL=DEBUG
GRAO_POLL_EVERY=@daily
RAO_POLL_FROM=/
BDC_POLL_FROM=/
AUTH0_AUDIENCE=https://grao.zenika.com/api/v1
AUTH0_DOMAIN=zenika.eu.auth0.com
AUTH0_JWKS_URI=https://zenika.eu.auth0.com/.well-known/jwks.json
AUTH0_ISSUER=https://zenika.eu.auth0.com/
DOCD_PORT=8888
DOCD_HOST=docd
 */
func show_env() {
	log.Info("APP_PORT=" + os.Getenv("GRAO_APP_PORT"))
	log.Info("DOCD_PORT=" + os.Getenv("DOCD_PORT"))
	log.Info("DOCD_HOST=" + os.Getenv("DOCD_HOST"))
	log.Info("BDC_POLL_FROM=" + os.Getenv("BDC_POLL_FROM"))
	log.Info("GRAO_DBX_CURSOR=" + os.Getenv("GRAO_DBX_CURSOR"))
	log.Info("GRAO_LOG_FILE=" + os.Getenv("GRAO_LOG_FILE"))
	log.Info("GRAO_DBX_ROOT=" + os.Getenv("GRAO_DBX_ROOT"))
	log.Info("GRAO_POLL_EVERY=" + os.Getenv("GRAO_POLL_EVERY"))
	log.Info("GRAO_LOG_LEVEL=" + os.Getenv("GRAO_LOG_LEVEL"))
}

func main() {
	/* INIT LOGGING */
	log.Init()
	defer log.Close()
	log.Info("Application started")
	show_env()
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
	auth := auth0.New(
		os.Getenv("AUTH0_JWKS_URI"),
		os.Getenv("AUTH0_ISSUER"),
		os.Getenv("AUTH0_AUDIENCE"),
	)
	auth0Middleware := auth.UserAuthenticatedMiddleware
	handler := auth0Middleware(c.Handler(r))
	http.ListenAndServe(":" + os.Getenv("GRAO_APP_PORT"), handler)
}
