package dropbox

import (
	"github.com/Zenika/RAO/auth"
	"github.com/Zenika/RAO/log"
	"github.com/stacktic/dropbox"
	"io/ioutil"
	"path/filepath"
	"regexp"
)

type DbxDocument struct {
	Title   string
	Path    string
	Mime    string
	Content string
	Client  string
	Region  string
	Mtime   dropbox.DBTime
	Bytes   int64
	Sum     string
}

type DbxCb func(bytes []byte, doc DbxDocument)

var db *dropbox.Dropbox = auth.RequireDropboxClient()
var filterPattern string = `(?i)^.+/_{1,2}clients(_|\s){1}(?P<Region>\w+)(/(?P<Client>[\w\s]+)(/.*))*`
var filter = regexp.MustCompile(filterPattern)

// Walk dropbox tree. leaves
// will be
func Walk(root string, fn DbxCb) {
	entry, err := db.Metadata(root, true, false, "", "", 0)
	log.Error(err, log.FATAL)
	contents := entry.Contents
	for _, e := range contents {
		process(e, fn)
	}
}

func process(e dropbox.Entry, fn DbxCb) {
	matches := filter.FindStringSubmatch(e.Path)
	if nil == matches {
		return
	}
	if !e.IsDir {
		region := matches[2]
		client := matches[4]
		size := e.Bytes
		modified := e.Modified
		bytes, _ := download(e.Path)
		doc := DbxDocument{
			Title:   filepath.Base(e.Path),
			Path:    filepath.Dir(e.Path),
			Mime:    e.MimeType,
			Content: "",
			Client:  client,
			Region:  region,
			Mtime:   modified,
			Bytes:   size,
			Sum:     "",
		}
		fn(bytes, doc)
		return
	}
	Walk(e.Path, fn)
}

func download(src string) ([]byte, int64) {
	resp, size, err := db.Download(src, "", 0)
	defer resp.Close()
	bytes, err := ioutil.ReadAll(resp)
	log.Error(err, log.ERROR)
	return bytes, size
}
