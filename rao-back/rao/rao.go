package rao

import (
	"github.com/Zenika/RAO/search/algolia"
  "github.com/Zenika/RAO/dropbox"
	"github.com/Zenika/RAO/search"
  "github.com/Zenika/RAO/utils"
  "github.com/Zenika/RAO/docd"
	"github.com/Zenika/RAO/log"
	"io/ioutil"
	"net/http"
	"os"
  "io"
)

var documents []dropbox.DbxDocument
var searchService = search.New(algolia.New())

func IndexAllDropBoxDocuments(w http.ResponseWriter, r *http.Request) {
	root := os.Getenv("RAO_DBX_ROOT")
	dropbox.Walk(root, func(res io.ReadCloser, doc dropbox.DbxDocument){
		buffer, err := ioutil.ReadAll(res)
		defer res.Close()
		log.Error(err, log.FATAL)
		content, _, err := docd.Convert(buffer, doc.Mime)
		log.Error(err, log.FATAL)
		doc.Content = string(content[:])
		if len(doc.Content) > 0 {
				doc.Sum = utils.Md5Sum(doc.Content)
				searchService.Store([]dropbox.DbxDocument{doc})
		}
	});
}
