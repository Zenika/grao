package rao

import "github.com/Zenika/rao/rao-back/document"

type RaoDocument struct {
	document.IDocument
	document.BusinessDocument
	document.FullTextDocument
}
