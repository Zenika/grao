package dropbox

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/Zenika/RAO/auth"
	"github.com/Zenika/RAO/document"
	"github.com/Zenika/RAO/log"
	"github.com/stacktic/dropbox"
)

var srcDir string = os.Getenv("RAO_POLL_FROM")

type Dropbox struct {
	client *dropbox.Dropbox
}

func New() *Dropbox {
	return &Dropbox{
		client: auth.RequireDropboxClient(),
	}
}

func (db Dropbox) Poll(root string, filter document.DocumentFilter, handler document.DocumentHandler) {
	cursor := db.lastCursor()
	db.writeCursor(db.delta(cursor, root, filter, handler))
}

func (db Dropbox) LongPoll(root string, filter document.DocumentFilter, handler document.DocumentHandler) {
	cursor := db.lastCursor()
	for {
		db.writeCursor(db.delta(cursor, root, filter, handler))
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

func (db Dropbox) delta(cursor string, root string, filter document.DocumentFilter, handler document.DocumentHandler) string {
	dp, err := db.client.Delta(cursor, root)
	log.Error(err, log.ERROR)
	cursor = dp.Cursor.Cursor
	for _, e := range dp.Entries {
		db.handleDeltaEntry(e, filter, handler)
	}
	if dp.HasMore {
		cursor = db.delta(cursor, root, filter, handler)
	}
	return cursor
}

func (db Dropbox) handleDeltaEntry(e dropbox.DeltaEntry, filter document.DocumentFilter, handler document.DocumentHandler) {
	if nil == e.Entry {
		log.Debug("nil entry")
		return
	}
	doc := db.createDocument(*e.Entry)
	if nil == doc {
		log.Debug("nil doc")
		return
	}
	if !filter(doc) {
		log.Debug("filtered")
		return
	}
	handler(doc)
}

func (db Dropbox) DownloadFile(doc document.IDocument) ([]byte, int64) {
	log.Debug("download")
	fullPath := fmt.Sprintf("%s/%s", doc.GetPath(), doc.GetTitle())
	resp, size, err := db.client.Download(fullPath, "", 0)
	log.Error(err, log.ERROR)
	defer resp.Close()
	bytes, err := ioutil.ReadAll(resp)
	log.Error(err, log.ERROR)
	return bytes, size
}

func (db Dropbox) createDocument(e dropbox.Entry) document.IDocument {
	if e.IsDir {
		log.Debug("dir found, return nil")
		return nil
	}
	modified := e.Modified
	doc := &document.Document{
		Title:     filepath.Base(e.Path),
		Path:      filepath.Dir(e.Path),
		Extension: filepath.Ext(e.Path),
		Mime:      e.MimeType,
		Mtime:     time.Time(modified),
	}
	log.Debug("document created")
	return doc
}

func (db Dropbox) lastCursor() string {
	b, err := ioutil.ReadFile(db.cursorFileName())
	if err != nil {
		return ""
	}
	return string(b)
}

func (db Dropbox) writeCursor(cursor string) {
	cursorFileName := db.cursorFileName()
	err := ioutil.WriteFile(cursorFileName, []byte(cursor), 0644)
	log.Error(err, log.FATAL)
}

func (db Dropbox) cursorFileName() string {
	cursorFileName := os.Getenv("RAO_DBX_CURSOR")
	if 0 == len(cursorFileName) {
		cursorFileName = "cursor"
	}
	return cursorFileName
}
