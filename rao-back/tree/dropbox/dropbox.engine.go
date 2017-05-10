// dropbox Package is a TreeEngune implementation
// that uses dropbox as a repository for documents
//
// see https://github.com/stacktic/dropbox/blob/master/dropbox.go
package dropbox

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/Zenika/RAO/auth"
	"github.com/Zenika/RAO/document"
	"github.com/Zenika/RAO/log"
	"github.com/stacktic/dropbox"
)

type Dropbox struct {
	client *dropbox.Dropbox
}

func New() *Dropbox {
	return &Dropbox{
		client: auth.RequireDropboxClient(),
	}
}

func (db Dropbox) Poll(root string, pairs [][]interface{}) {
	cursorFileName := db.cursorFileName()
	cursor := db.lastCursor(cursorFileName)
	db.writeCursor(db.delta(cursor, root, pairs), cursorFileName)
}

func (db Dropbox) LongPoll(root string, pairs [][]interface{}) {
	cursorFileName := db.cursorFileName()
	cursor := db.lastCursor(cursorFileName)
	for {
		db.writeCursor(db.delta(cursor, root, pairs), cursorFileName)
		changes := false
		for !changes {
			poll, err := db.client.LongPollDelta(cursor, 30)
			if err != nil {
				log.Error(err, log.ERROR)
			}
			changes = poll.Changes
			duration, _ := time.ParseDuration(strconv.Itoa(poll.Backoff) + "s")
			time.Sleep(duration)
		}
	}
}

func (db Dropbox) DownloadFile(doc document.IDocument) ([]byte, int64) {
	fullPath := fmt.Sprintf("%s/%s", doc.GetPath(), doc.GetTitle())
	resp, size, err := db.client.Download(fullPath, "", 0)
	log.Error(err, log.ERROR)
	defer resp.Close()
	bytes, err := ioutil.ReadAll(resp)
	log.Error(err, log.ERROR)
	return bytes, size
}

func (db Dropbox) delta(cursor string, root string, pairs [][]interface{}) string {
	dp, err := db.client.Delta(cursor, root)
	log.Error(err, log.ERROR)
	cursor = dp.Cursor.Cursor
	log.Debug("poll " + root)
	for _, e := range dp.Entries {
		for _, p := range pairs {
			db.handleDeltaEntry(e, p[0].(func(document.IDocument) bool), p[1].(func(document.IDocument)))
		}
		if dp.HasMore {
			cursor = db.delta(cursor, root, pairs)
		}
	}
	return cursor
}

func (db Dropbox) handleDeltaEntry(e dropbox.DeltaEntry, filter func(document.IDocument) bool, handler func(document.IDocument)) {
	if nil == e.Entry {
		return
	}
	doc := db.createDocument(*e.Entry)
	if nil == doc {
		return
	}
	if !filter(doc) {
		return
	}
	log.Debug("handle " + doc.GetPath())
	handler(doc)
}

func (db Dropbox) createDocument(e dropbox.Entry) document.IDocument {
	if e.IsDir {
		return nil
	}
	modified := e.Modified
	doc := &document.Document{
		Title:     filepath.Base(e.Path),
		Path:      filepath.Dir(e.Path),
		Extension: strings.TrimPrefix(filepath.Ext(e.Path), "."),
		Mime:      e.MimeType,
		Mtime:     time.Time(modified),
	}
	return doc
}

func (db Dropbox) lastCursor(filename string) string {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return ""
	}
	return string(b)
}

func (db Dropbox) writeCursor(cursor string, filename string) {
	err := ioutil.WriteFile(filename, []byte(cursor), 0644)
	log.Error(err, log.FATAL)
}

func (db Dropbox) cursorFileName() string {
	cursorFileName := os.Getenv("GRAO_DBX_CURSOR")
	if 0 == len(cursorFileName) {
		cursorFileName = "cursor"
	}
	return cursorFileName
}
