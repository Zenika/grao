package controller

import (
	"fmt"
	"os"
	"github.com/Zenika/rao/rao-back/log"
	"net/http"

	"github.com/Zenika/rao/rao-back/conv"
	bdcService "github.com/Zenika/rao/rao-back/document/bdc/service"
	raoService "github.com/Zenika/rao/rao-back/document/rao/service"
	"github.com/Zenika/rao/rao-back/search"
	"github.com/Zenika/rao/rao-back/tree"
)

func GrabHandler(searchService *search.SearchService, convService *conv.ConvService, treeService *tree.TreeService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Debug("GrabHandler method called")
		grabAndConvertDocuments(searchService, convService, treeService)
		handleGrab(w, r)
	}
}

func grabAndConvertDocuments(searchService *search.SearchService, convService *conv.ConvService, treeService *tree.TreeService) {
	log.Debug("grabAndConvertDocuments method called")
	root := fmt.Sprintf("/%v", os.Getenv("GRAO_DBX_ROOT"))
	bdcService := bdcService.New(*searchService, *treeService)
	raoService := raoService.New(*searchService, *convService, *treeService)
	pairs := [][]interface{}{{bdcService.DocFilter, bdcService.DocHandler}, {raoService.DocFilter, raoService.DocHandler}}
	treeService.Poll(root, pairs)
}

func handleGrab(w http.ResponseWriter, r *http.Request) {
	log.Debug("handleGrab method called")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Retrieving and converting documents from Dropbox")))
}