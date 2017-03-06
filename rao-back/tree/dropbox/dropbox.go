package dropbox

import (
	"github.com/Zenika/RAO/auth"
	"github.com/Zenika/RAO/log"
	"github.com/Zenika/RAO/document"
	"github.com/stacktic/dropbox"
  "strconv"
	"io/ioutil"
	"time"
	"path/filepath"
	"regexp"
)

// var db *dropbox.Dropbox = auth.RequireDropboxClient()
var filterPattern string = `(?i)^.+/_{1,2}clients(_|\s){1}(?P<Region>\w+)(/(?P<Client>[\w\s]+)(/.*))*`
var filter = regexp.MustCompile(filterPattern)


type Dropbox struct {
	client *dropbox.Dropbox
}

func New() *Dropbox {
	return &Dropbox{
		client: auth.RequireDropboxClient(),
	}
}

func (db Dropbox) Walk(root string, handler document.DocumentHandler) {
	entry, err := db.client.Metadata(root, true, false, "", "", 0)
	log.Error(err, log.FATAL)
	contents := entry.Contents
	for _, e := range contents {
    matches := filter.FindStringSubmatch(e.Path)
    if nil == matches {
      continue
    }
    if !e.IsDir {
      // region := matches[2]
      // client := matches[4]
      // size := e.Bytes
      // modified := e.Modified
      // bytes, _ := db.download(e.Path)
      // doc := &document.Document{
      //   Title:   filepath.Base(e.Path),
      //   Path:    filepath.Dir(e.Path),
      //   Mime:    e.MimeType,
      //   Content: "",
      //   Client:  client,
      //   Region:  region,
      //   Mtime:   time.Time(modified),
      //   Bytes:   size,
      //   Sum:     "",
      // }
      doc := db.createDocument(e)
      bytes, _ := db.downloadFile(e)
      handler(bytes, doc)
      return
    }
    db.Walk(e.Path, handler)
	}
}

func (db Dropbox) Poll(root string, handler document.DocumentHandler){
  cursor := ""
  for {
    dp, err := db.client.Delta(cursor, root)
    log.Error(err, log.ERROR)
    for _, e := range dp.Entries {
      log.Debug(e.Entry.Path)
      // doc := db.createDocument(*e.Entry)
      // bytes, _ := db.downloadFile(*e.Entry)
    }
    cursor = dp.Cursor.Cursor
    changes := false
    for !changes {
      poll, _ := db.client.LongPollDelta(cursor, 30)
      changes = poll.Changes
      log.Debug("backoff: " + strconv.Itoa(poll.Backoff) + "s")
      duration, _ := time.ParseDuration(strconv.Itoa(poll.Backoff) + "s")
      time.Sleep(duration)
    }
  }
}

func (db Dropbox) downloadFile(e dropbox.Entry) ([]byte, int64) {
	resp, size, err := db.client.Download(e.Path, "", 0)
	defer resp.Close()
	bytes, err := ioutil.ReadAll(resp)
	log.Error(err, log.ERROR)
	return bytes, size
}

func (db Dropbox) createDocument(e dropbox.Entry) (document.IDocument){
  matches := filter.FindStringSubmatch(e.Path)
  if nil == matches {
    return nil
  }
  if !e.IsDir {
    region := matches[2]
    client := matches[4]
    size := e.Bytes
    modified := e.Modified
    doc := &document.Document{
      Title:   filepath.Base(e.Path),
      Path:    filepath.Dir(e.Path),
      Mime:    e.MimeType,
      Content: "",
      Client:  client,
      Region:  region,
      Mtime:   time.Time(modified),
      Bytes:   size,
      Sum:     "",
    }
    return doc
  }
  return nil
}
