package rao

import "github.com/Zenika/RAO/document"

type RaoDocument struct {
	document.IDocument
	document.BusinessDocument
	document.FullTextDocument
}
