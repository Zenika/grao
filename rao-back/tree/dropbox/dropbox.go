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
// "application/vnd.openxmlformats-officedocument.wordprocessingml.document"
var mimes = []string{"application/pdf"}

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
		if !e.IsDir {
			doc := db.createDocument(e)
			if nil == doc {
				continue
			}
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

func (db Dropbox) delta(cursor string, root string, handler document.DocumentHandler) string {
	log.Debug("cursor " + cursor)
	dp, err := db.client.Delta(cursor, root)
	log.Error(err, log.ERROR)
	cursor = dp.Cursor.Cursor
	db.writeCursor(cursor)
	for _, e := range dp.Entries {
		db.handleDeltaEntry(e, handler)
	}
	if dp.HasMore {
		db.delta(cursor, root, handler)
	}
	return cursor
}

func (db Dropbox) handleDeltaEntry(e dropbox.DeltaEntry, handler document.DocumentHandler) {
	if nil == e.Entry {
		log.Debug("nil entry")
		return
	}
	doc := db.createDocument(*e.Entry)
	if nil == doc {
		log.Debug("nil doc")
		return
	}
	bytes, _ := db.downloadFile(*e.Entry)
	handler(bytes, doc)
}

func (db Dropbox) downloadFile(e dropbox.Entry) ([]byte, int64) {
	resp, size, err := db.client.Download(e.Path, "", 0)
	log.Error(err, log.ERROR)
	defer resp.Close()
	bytes, err := ioutil.ReadAll(resp)
	log.Error(err, log.ERROR)
	return bytes, size
}

func (db Dropbox) createDocument(e dropbox.Entry) document.IDocument {
	log.Debug(e.Path)
	if e.IsDir {
		return nil
	}
	if !utils.ArrayContainsString(mimes, e.MimeType) {
		return nil
	}
	matches := filter.FindStringSubmatch(e.Path)
	if nil == matches {
		log.Debug("no match")
		return nil
	}
	agence := matches[2]
	client := matches[4]
	size := e.Bytes
	modified := e.Modified
	doc := &document.Document{
		Title:   				filepath.Base(e.Path),
		Path:    				filepath.Dir(e.Path),
		Ext:     				filepath.Ext(e.Path),
		Mime:    				e.MimeType,
		Content: 				"",
		Client:  				client,
		Agence:  				agence,
		Mtime:   				time.Time(modified),
		Bytes:   				size,
		Sum:     				"",
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
