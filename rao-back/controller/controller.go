package controller

import (
	"encoding/json"
	"fmt"
	"github.com/Zenika/RAO/conv"
	"github.com/Zenika/RAO/conv/docd"
	"github.com/Zenika/RAO/document"
	"github.com/Zenika/RAO/log"
	"github.com/Zenika/RAO/search"
	"github.com/Zenika/RAO/search/algolia"
	"github.com/Zenika/RAO/tree"
	"github.com/Zenika/RAO/tree/dropbox"
	"github.com/Zenika/RAO/utils"
	"net/http"
	"os"
)

var searchService = search.New(algolia.New())
var convService = conv.New(docd.New())
var treeService = tree.New(dropbox.New())

func Search(w http.ResponseWriter, r *http.Request) {
	var query search.SearchQuery
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&query)
	if nil != err {
		log.Error(err, log.ERROR)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("There was an error: %v", err)))
		return
	}
	res, err := searchService.Search(query)
	if err == nil {
		w.Write([]byte(res))
	} else {
		log.Error(err, log.ERROR)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("There was an error: %v", err)))
	}
}

func Walk(w http.ResponseWriter, r *http.Request) {
	root := os.Getenv("RAO_DBX_ROOT")
	log.Debug("root " + root)
	treeService.Walk(root, func(bytes []byte, doc document.IDocument) {
		b, err := convService.Convert(bytes, doc.GetMime())
		content := string(b[:])
		log.Error(err, log.ERROR)
		chunks := utils.SplitString(content, 10000)
		doc.SetSum(utils.Md5Sum(content))
		for _, chunk := range chunks {
			doc.SetContent(chunk)
			searchService.Store([]document.IDocument{doc})
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
			searchService.Store([]document.IDocument{doc})
		}
	})
}

func Poll() {
	root := fmt.Sprintf("/%v", os.Getenv("RAO_DBX_ROOT"))
	treeService.Poll(root, func(bytes []byte, doc document.IDocument) {
		b, err := convService.Convert(bytes, doc.GetMime())
		content := string(b[:])
		log.Debug(doc.GetPath())
		if len(content) == 0 {
			return
		}
		log.Error(err, log.ERROR)
		chunks := utils.SplitString(content, 10000)
		doc.SetSum(utils.Md5Sum(content))
		for _, chunk := range chunks {
			doc.SetContent(chunk)
			searchService.Store([]document.IDocument{doc})
		}
	})
}
