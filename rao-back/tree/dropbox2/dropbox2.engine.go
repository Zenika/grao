// Test for dropbox-sdk-unofficial
// see https://github.com/dropbox/dropbox-sdk-go-unofficial.git
package dropbox2

import (
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/files"
	"github.com/Zenika/rao/rao-back/auth"
	"os"
	"github.com/Zenika/rao/rao-back/log"
	"github.com/Zenika/rao/rao-back/document"
	"path/filepath"
)

type Dropbox2 struct {
	client files.Client
}

func New() *Dropbox2 {
	return &Dropbox2{
		client: auth.RequireDropbox2Client(),
	}
}

func (db Dropbox2) cursorFileName() string {
	cursorFileName := os.Getenv("GRAO_DBX_CURSOR")
	if 0 == len(cursorFileName) {
		cursorFileName = "cursor"
	}
	return cursorFileName
}

func (db Dropbox2) delta(cursor string, root string, pairs  [][]interface{}) string {
	arg := *files.NewListFolderArg("")
	arg.Recursive = true
	// Getting list of Metadatas
	dp, err :=  db.client.ListFolder(&arg)
	log.Error(err, log.ERROR)
	cursor = dp.Cursor
	log.Debug("Polling : " + root)
	for _, entry := range dp.Entries {
		switch entry.(type) {
		case *files.FileMetadata:
			fileEntry, _ := entry.(*files.FileMetadata)
			db.createDocument(*fileEntry)


		}
	}
	return cursor
}

func (db Dropbox2) createDocument(entry files.FileMetadata) document.IDocument {
	modified := entry.ClientModified
	doc := &document.Document{
		Title: filepath.Base(entry.PathLower),
		Path: filepath.Dir(entry.PathLower),
		Mtime: modified,
	}
	return doc
}