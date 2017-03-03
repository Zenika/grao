package dropbox

import (
	"github.com/stacktic/dropbox"
	"github.com/Zenika/RAO/auth"
	"github.com/Zenika/RAO/log"
	"path/filepath"
	"regexp"
	"fmt"
	"io"
)

type DbxDocument struct {
	Title   string
	Path    string
	Mime    string
	Content string
	Client  string
	Region  string
	Mtime		dropbox.DBTime
	Bytes   int64
  Sum     string
}

type DbxCb func(res io.ReadCloser, doc DbxDocument)

var db *dropbox.Dropbox = auth.RequireDropboxClient()
var filterPattern string = `(?i)^.+/_{1,2}clients(_|\s){1}(?P<Region>\w+)(/(?P<Client>[\w\s]+)(/.*))*`
var filter = regexp.MustCompile(filterPattern)

func Walk(root string, fn DbxCb) {
	entry, err := db.Metadata(root, true, false, "", "", 0)
	log.Error(err, log.FATAL)
	contents := entry.Contents
	for _, e := range contents {
    process(e, fn)
	}
}

func process(e dropbox.Entry, fn DbxCb){
	matches := filter.FindStringSubmatch(e.Path)
	if nil == matches {
		return
	}
  if !e.IsDir {
		log.Debug(fmt.Sprintf("%v",matches));
		region := matches[2]
		client := matches[4]
		bytes := e.Bytes
		modified := e.Modified
    res, _ := download(e.Path)
		doc := DbxDocument {
			Title: filepath.Base(e.Path),
			Path: e.Path,
			Mime: e.MimeType,
			Content: "",
			Client: client,
			Region: region,
			Mtime: modified,
			Bytes: bytes,
      Sum: "",
		}
    fn(res, doc)
    return
  }
  Walk(e.Path, fn)
}


func download(src string) (io.ReadCloser, int64) {
	reader, size, err := db.Download(src, "", 0)
	log.Error(err, log.ERROR)
	return reader, size
}
