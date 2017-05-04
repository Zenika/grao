package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Zenika/RAO/conv"
	"github.com/Zenika/RAO/conv/docd"
	"github.com/Zenika/RAO/document"
	"github.com/Zenika/RAO/log"
	"github.com/Zenika/RAO/search"
	"github.com/Zenika/RAO/search/algolia"
	searchController "github.com/Zenika/RAO/search/controller"
	"github.com/Zenika/RAO/tree"
	"github.com/Zenika/RAO/tree/dropbox"
	"github.com/Zenika/RAO/utils"
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
	cronExp := os.Getenv("RAO_POLL_EVERY")
	if len(cronExp) == 0 {
		cronExp = "@daily" // equivalent to 0 0 0 * * *
	}
	cron := cron.New()
	cron.AddFunc(cronExp, func() {
		root := fmt.Sprintf("/%v", os.Getenv("RAO_DBX_ROOT"))
		treeService.Poll(root, func(bytes []byte, doc document.IDocument) {
			b, err := convService.Convert(bytes, doc.GetMime())
			content := string(b[:])
			if len(content) == 0 {
				return
			}
			log.Error(err, log.ERROR)
			chunks := utils.SplitString(content, 10000)
			doc.SetSum(utils.Md5Sum(content))
			for _, chunk := range chunks {
				doc.SetContent(chunk)
				searchService.Store([]document.IDocument{doc}, algolia.SearchOptions{Index: "rao"})
			}
		})
	})
	cron.Start()
	/* INIT HTTP CONTROLLER */
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "DELETE", "OPTION", "PUT"},
		AllowCredentials: true,
	})
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/{index}", searchController.SearchHandler(searchService)).Methods("POST")
	handler := c.Handler(r)
	http.ListenAndServe(":8090", handler)
}
