package dropbox

import (
	"github.com/Zenika/RAO/auth"
	"github.com/Zenika/RAO/document"
	"github.com/Zenika/RAO/log"
	"github.com/Zenika/RAO/utils"
	"github.com/stacktic/dropbox"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"time"
)

var filterPattern string = `(?i)^.+/_{1,2}clients(_|\s){1}(?P<Agence>[\w&\s]+)(/(?P<Client>[\w\s]+)(/.*))*`
var filter = regexp.MustCompile(filterPattern)
var mimes = []string{"application/pdf", "application/vnd.openxmlformats-officedocument.wordprocessingml.document"}

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
			doc := db.createDocument(e)
			bytes, _ := db.downloadFile(e)
			handler(bytes, doc)
			return
		}
		db.Walk(e.Path, handler)
	}
}

func (db Dropbox) Poll(root string, handler document.DocumentHandler) {
	cursor := db.lastCursor()
	for {
		cursor = db.delta(cursor, root, handler)
		changes := false
		for !changes {
			poll, _ := db.client.LongPollDelta(cursor, 30)
			changes = poll.Changes
			duration, _ := time.ParseDuration(strconv.Itoa(poll.Backoff) + "s")
			time.Sleep(duration)
		}
	}
}

func (db Dropbox) delta(cursor string, root string, handler document.DocumentHandler) string {
	dp, err := db.client.Delta(cursor, root)
	db.writeCursor(cursor)
	log.Error(err, log.ERROR)
	cursor = dp.Cursor.Cursor
	for _, e := range dp.Entries {
		db.handleDeltaEntry(e, handler)
	}
	if dp.HasMore {
		db.delta(cursor, root, handler)
	}
	return cursor
}

func (db Dropbox) handleDeltaEntry(e dropbox.DeltaEntry, handler document.DocumentHandler) {
	doc := db.createDocument(*e.Entry)
	if nil == doc {
		return
	}
	bytes, _ := db.downloadFile(*e.Entry)
	handler(bytes, doc)
}

func (db Dropbox) downloadFile(e dropbox.Entry) ([]byte, int64) {
	resp, size, err := db.client.Download(e.Path, "", 0)
	defer resp.Close()
	bytes, err := ioutil.ReadAll(resp)
	log.Error(err, log.ERROR)
	return bytes, size
}

func (db Dropbox) createDocument(e dropbox.Entry) document.IDocument {
	if e.IsDir {
		return nil
	}
	if !utils.ArrayContainsString(mimes, e.MimeType) {
		return nil
	}
	matches := filter.FindStringSubmatch(e.Path)
	if nil == matches {
		return nil
	}
	agence := matches[2]
	client := matches[4]
	size := e.Bytes
	modified := e.Modified
	doc := &document.Document{
		Title:   filepath.Base(e.Path),
		Path:    filepath.Dir(e.Path),
		Ext:     filepath.Ext(e.Path),
		Mime:    e.MimeType,
		Content: "",
		Client:  client,
		Agence:  agence,
		Mtime:   time.Time(modified),
		Bytes:   size,
		Sum:     "",
	}
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
	cursorFileName := os.Getenv("DBX_CURSOR_FILE")
	if 0 == len(cursorFileName) {
		cursorFileName = "cursor"
	}
	return cursorFileName
}
