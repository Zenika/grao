package rao

import (
	"github.com/Zenika/RAO/conv"
	"github.com/Zenika/RAO/conv/docd"
	"github.com/Zenika/RAO/dropbox"
	"github.com/Zenika/RAO/log"
	"github.com/Zenika/RAO/search"
	"github.com/Zenika/RAO/search/algolia"
	"github.com/Zenika/RAO/utils"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

var documents []dropbox.DbxDocument

var searchService = search.New(algolia.New())
var convService = conv.New(docd.New())

func IndexAllDropBoxDocuments(w http.ResponseWriter, r *http.Request) {
	root := os.Getenv("RAO_DBX_ROOT")
	dropbox.Walk(root, func(res io.ReadCloser, doc dropbox.DbxDocument) {
		buffer, err := ioutil.ReadAll(res)
		defer res.Close()
		log.Error(err, log.FATAL)
		b, err := convService.Convert(buffer, doc.Mime)
		content := string(b[:])
		log.Error(err, log.ERROR)
		chunks := utils.SplitString(content, 10000)
		doc.Sum = utils.Md5Sum(content)
		for _, chunk := range chunks {
			doc.Content = chunk
			searchService.Store([]dropbox.DbxDocument{doc})
		}
	})
}

func Search(w http.ResponseWriter, r *http.Request) {
	pattern := r.URL.Query().Get("query")
	res, err := searchService.Search(pattern)
	if err == nil {
		w.Write([]byte(res))
	} else {
		log.Error(err, log.ERROR)
	}
}
