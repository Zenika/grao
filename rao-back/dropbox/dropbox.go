package dropbox

import (
	"github.com/Zenika/RAO/auth"
	"github.com/Zenika/RAO/docd"
	"github.com/stacktic/dropbox"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
  "path/filepath"
)

type DbxDocument struct {
	Title   string
	Path    string
	Mime    string
	Content string
}

type DbxCb func(res io.ReadCloser, doc DbxDocument)

var db *dropbox.Dropbox = auth.RequireDropboxClient()

func GetRootFolder(w http.ResponseWriter, r *http.Request) {
	root := os.Getenv("RAO_DBX_ROOT")
	Walk(root, func(res io.ReadCloser, doc DbxDocument){
    content := convert(res, doc.Mime)
    doc.Content = content
    log.Println(doc.Content)
  });
}

func Walk(root string, fn DbxCb) {
	entry, err := db.Metadata(root, true, false, "", "", 0)
	check(err)
	contents := entry.Contents
	for _, e := range contents {
    process(e, fn)
	}
}

func process(e dropbox.Entry, fn DbxCb){
  if !e.IsDir {
    res, _ := download(e.Path)
    fn(res, DbxDocument { filepath.Base(e.Path), e.Path, e.MimeType, ""})
    return
  }
  Walk(e.Path, fn)
}

func convert(rc io.ReadCloser, mime string) string {
	buffer, err := ioutil.ReadAll(rc)
	defer rc.Close()
	check(err)
	body, _, err := docd.Convert(buffer, mime)
	check(err)
	return string(body[:])
}

func download(src string) (io.ReadCloser, int64) {
	reader, size, err := db.Download(src, "", 0)
	check(err)
	return reader, size
}


func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
