package refs

import "github.com/Zenika/RAO/document"

type RefsDocument struct {
	document.IDocument
	document.ReferencesDocument
	document.FullTextDocument
}
