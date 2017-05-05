package main

import (
	"fmt"
	"net/http"
	"os"
	"regexp"

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
	"github.com/algolia/algoliasearch-client-go/algoliasearch"
	"github.com/gorilla/mux"
	"github.com/robfig/cron"
	"github.com/rs/cors"
)

/* INIT SERVICES IMPLEMENTATIONS */
var treeService = tree.New(dropbox.New())
var convService = conv.New(docd.New())
var searchService = search.New(algolia.New())

var srcDir string = os.Getenv("RAO_POLL_FROM")

var raoFilterPattern string = fmt.Sprintf(
	`(?i)^.+/_{1,2}clients(_|\s){1}(?P<Agence>[\w&\s]+)/(?P<Client>[^/]+)/%s`,
	srcDir)
var raoPatternFilter = regexp.MustCompile(raoFilterPattern)

// Adding support for docx documents:
// "application/vnd.openxmlformats-officedocument.wordprocessingml.document"
var mimes = []string{"application/pdf"}

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

	raoFilter := func(doc document.IDocument) bool {
		if !utils.ArrayContainsString(mimes, doc.GetMime()) {
			log.Debug("bad mime " + doc.GetMime())
			return false
		}
		matches := raoPatternFilter.FindStringSubmatch(doc.GetPath())
		if nil == matches {
			log.Debug("no match " + doc.GetPath())
			return false
		}
		return true
	}

	docMapper := func(doc document.IDocument) interface{} {
		return algoliasearch.Object{
			"Content":   doc.(document.RaoDocument).GetContent(),
			"Title":     doc.GetTitle(),
			"Path":      doc.GetPath(),
			"Client":    doc.(document.RaoDocument).GetClient(),
			"Agence":    doc.(document.RaoDocument).GetAgence(),
			"Extension": doc.GetExtension(),
			"Mime":      doc.GetMime(),
			"Mtime":     doc.GetMtime(),
			"Bytes":     doc.(document.RaoDocument).GetBytes(),
			"Sum":       doc.(document.RaoDocument).GetSum(),
		}
	}

	raoHandler := func(doc document.IDocument) {
		bytes, size := treeService.GetEngine().(*dropbox.Dropbox).DownloadFile(doc)
		b, err := convService.Convert(bytes, doc.GetMime())
		log.Error(err, log.ERROR)
		content := string(b[:])
		if "" == content {
			return // Shall we index the document if we could not extract its content ?
		}
		matches := raoPatternFilter.FindStringSubmatch(doc.GetPath())
		agence := matches[2]
		client := matches[3]
		chunks := utils.SplitString(content, 10000)
		for _, chunk := range chunks {
			raoDocument := document.RaoDocument{
				doc,
				document.BusinessDocument{
					doc,
					agence,
					client,
				},
				document.FullTextDocument{
					Bytes:   size,
					Sum:     utils.Md5Sum(content),
					Content: chunk,
				},
			}
			searchService.Store(raoDocument, docMapper, algolia.SearchOptions{Index: "rao"})
		}
	}
	cron.AddFunc(cronExp, func() {
		root := fmt.Sprintf("/%v", os.Getenv("RAO_DBX_ROOT"))
		treeService.Poll(root, raoFilter, raoHandler)
	})
	cron.Start()
	/* INIT HTTP CONTROLLER */
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "DELETE", "OPTION", "PUT"},
		AllowCredentials: true,
	})
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/{index}/search", searchController.SearchHandler(searchService)).Methods("POST")
	r.HandleFunc("/api/v1/{index}/settings", searchController.SettingsHandler(searchService)).Methods("POST")
	handler := c.Handler(r)
	http.ListenAndServe(":8090", handler)
}
