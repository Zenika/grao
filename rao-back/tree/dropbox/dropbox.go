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
	"github.com/Zenika/RAO/utils"
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

func (db Dropbox) Poll(root string, pairs [][]interface{}) {
	cursor := db.lastCursor(utils.Md5Sum(root))
	db.writeCursor(db.delta(cursor, root, pairs), utils.Md5Sum(root))
}

func (db Dropbox) LongPoll(root string, pairs [][]interface{}) {
	cursor := db.lastCursor(utils.Md5Sum(root))
	for {
		db.writeCursor(db.delta(cursor, root, pairs), utils.Md5Sum(root))
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

func (db Dropbox) delta(cursor string, root string, pairs [][]interface{}) string {
	dp, err := db.client.Delta(cursor, root)
	log.Debug("dropbox " + root)
	log.Error(err, log.ERROR)
	cursor = dp.Cursor.Cursor
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
	handler(doc)
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

func (db Dropbox) createDocument(e dropbox.Entry) document.IDocument {
	if e.IsDir {
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
	// cursorFileName := db.cursorFileName()
	err := ioutil.WriteFile(filename, []byte(cursor), 0644)
	log.Error(err, log.FATAL)
}

func (db Dropbox) cursorFileName() string {
	cursorFileName := os.Getenv("RAO_DBX_CURSOR")
	if 0 == len(cursorFileName) {
		cursorFileName = "cursor"
	}
	return cursorFileName
}
