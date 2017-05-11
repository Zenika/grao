package tree

import (
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/Zenika/RAO/document"
)

type Mock struct {
	docs []document.IDocument
}

func (mock Mock) Poll(root string, pairs [][]interface{}) {

}

func (mock Mock) LongPoll(root string, pairs [][]interface{}) {

}

var treeEngine = &mock{
	docs: []document.IDocument{
		createDoc("Zenika Clients/_clients_lille/lex luthor/krypto project/REPONSE ENVOYEE/response.pdf"),
		createDoc("Zenika Clients/_clients_lille/Kahuna Burger/bruger online/REPONSE ENVOYEE/response.pdf"),
	},
}

func createDoc(path string) document.IDocument {
	return &document.Document{
		Title:     filepath.Base(path),
		Path:      filepath.Dir(path),
		Extension: strings.TrimPrefix(filepath.Ext(path), "."),
		Mime:      "application/pdf",
		Mtime:     time.Now(),
	}
}

func Test_Poll(t *testing.T) {
	t.Error("not implemented")
}

func Test_LongPoll(t *testing.T) {
	t.Error("not implemented")
}
