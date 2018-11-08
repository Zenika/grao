// Test for dropbox-sdk-unofficial
// see https://github.com/dropbox/dropbox-sdk-go-unofficial.git
package dropbox2

import (
	"github.com/Zenika/rao/rao-back/auth"
	"github.com/Zenika/rao/rao-back/document"
	"github.com/Zenika/rao/rao-back/log"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/files"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"fmt"
)

var REFERER = log.GetReferer()

type Dropbox2 struct {
	client files.Client
}

func New() *Dropbox2 {
	return &Dropbox2{
		client: auth.RequireDropbox2Client(),
	}
}

func (db Dropbox2) LongPoll(root string, pairs [][]interface{}) {
	panic("implement me")
}

func (db Dropbox2) Poll(root string, pairs [][]interface{}) {
	cursorFileName := db.cursorFileName()
	cursor := db.lastCursor(cursorFileName)
	db.writeCursor(db.delta(cursor, root, pairs), cursorFileName)
}

func (db Dropbox2) DownloadFile(doc document.IDocument) ([]byte, int64) {
	fullPath := fmt.Sprintf("%s/%s", doc.GetPath(), doc.GetTitle())
	downloadArg := files.NewDownloadArg(fullPath)
	metadata, resp, err := db.client.Download(downloadArg)
	if err != nil {
		log.Error(err, log.ERROR, REFERER) // TODO: implement clearer errors
	}
	bytes, err := ioutil.ReadAll(resp)
	if err != nil {
		log.Error(err, log.ERROR, REFERER)
	}
	size := int64(metadata.Size)
	return bytes, size
}

func (db Dropbox2) delta(cursor string, root string, pairs [][]interface{}) string {
	arg := *files.NewListFolderArg("")
	arg.Recursive = true
	// Getting the list of Entries
	listFolderResult, err := db.client.ListFolder(&arg)
	log.Error(err, log.ERROR, REFERER)
	cursor = listFolderResult.Cursor
	log.Debug("Polling : " + root, REFERER)
	// Iterating over Entries and casting types to access the data and map objects
	for _, entry := range listFolderResult.Entries {
		for _, p := range pairs {
			db.handleDeltaEntry(entry, p[0].(func(document.IDocument) bool), p[1].(func(document.IDocument)))
		}
		// Calls again if more results to handle
		if listFolderResult.HasMore {
			log.Debug("HasMore", REFERER)
			//cursor = db.delta(cursor, root, pairs)
		}
	}
	return cursor
}

// Handles the files metadatas, creates doc, filters and index if needed
func (db Dropbox2) handleDeltaEntry(metadata files.IsMetadata, filter func(document.IDocument) bool, handler func(document.IDocument)) {
	switch metadata.(type) {
	// Casting types to access the metadata
	case *files.FileMetadata:
		// Handles case of file, cast type and create doc
		fileEntry, _ := metadata.(*files.FileMetadata)
		log.Debug("Document found :" + fileEntry.PathLower, REFERER)
		doc := db.createDocument(*fileEntry)
		// Pass document to the handler function for indexing
		if filter(doc) {
			log.Debug("Handling " + doc.GetPath(), REFERER)
			handler(doc)
		}
	//case *files.FolderMetadata:
	//	// Handles case of folder, maybe useless... apart for log
	//	folderEntry, _ := metadata.(*files.FolderMetadata)
	//	log.Debug("Folder found : " + folderEntry.PathLower)
	//	return
	default:
		log.Debug("None found", REFERER)
		return
	}

}

// Handle the mapping of Filemetadata to Document
func (db Dropbox2) createDocument(entry files.FileMetadata) document.IDocument {
	modified := entry.ServerModified
	extension := strings.TrimPrefix(filepath.Ext(entry.PathDisplay), ".")
	doc := &document.Document{
		Title:     filepath.Base(entry.PathDisplay),
		Path:      filepath.Dir(entry.PathDisplay),
		Extension: extension,
		Mtime:     modified,
		Mime:      "application/pdf", // TODO : implements other MimeTypes
	}
	return doc
}

func (db Dropbox2) lastCursor(filename string) string {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return ""
	}
	return string(b)
}

func (db Dropbox2) writeCursor(cursor string, filename string) {
	err := ioutil.WriteFile(filename, []byte(cursor), 0644)
	log.Error(err, log.FATAL, REFERER)
}

func (db Dropbox2) cursorFileName() string {
	cursorFileName := os.Getenv("GRAO_DBX_CURSOR")
	if 0 == len(cursorFileName) {
		cursorFileName = "cursor"
	}
	return cursorFileName
}