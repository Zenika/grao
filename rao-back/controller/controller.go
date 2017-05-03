package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/Zenika/RAO/conv"
	"github.com/Zenika/RAO/conv/docd"
	"github.com/Zenika/RAO/document"
	"github.com/Zenika/RAO/log"
	"github.com/Zenika/RAO/search"
	"github.com/Zenika/RAO/search/algolia"
	"github.com/Zenika/RAO/tree"
	"github.com/Zenika/RAO/tree/dropbox"
	"github.com/Zenika/RAO/utils"
)

var searchService = search.New(algolia.New())
var convService = conv.New(docd.New())
var treeService = tree.New(dropbox.New())

func Search(w http.ResponseWriter, r *http.Request) {
	var query search.Query
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&query)
	if nil != err {
		log.Error(err, log.ERROR)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("There was an error: %v", err)))
		return
	}
	queryRes, err := searchService.Search(query, algolia.SearchOptions{Index: "rao"})
	response, err := json.Marshal(queryRes.Data)
	if err == nil {
		w.Write([]byte(response))
	} else {
		log.Error(err, log.ERROR)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("There was an error: %v", err)))
	}
}

func Walk(w http.ResponseWriter, r *http.Request) {
	root := os.Getenv("RAO_DBX_ROOT")
	treeService.Walk(root, func(bytes []byte, doc document.IDocument) {
		b, err := convService.Convert(bytes, doc.GetMime())
		content := string(b[:])
		log.Error(err, log.ERROR)
		chunks := utils.SplitString(content, 10000)
		doc.SetSum(utils.Md5Sum(content))
		for _, chunk := range chunks {
			doc.SetContent(chunk)
			searchService.Store([]document.IDocument{doc}, algolia.SearchOptions{Index: "rao"})
		}
	})
}

func LongPoll() {
	root := fmt.Sprintf("/%v", os.Getenv("RAO_DBX_ROOT"))
	go treeService.LongPoll(root, func(bytes []byte, doc document.IDocument) {
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
}

func Poll() {
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
}
