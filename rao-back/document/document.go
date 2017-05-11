// Package document contains common document interfaces,
// with subpackages related to document implementations
// and their associated services
package document

import "time"

// DocumentMapper is a mapper function that can be used
// by external services to convert any implentation of IDocument
// into a map
type DocumentMapper func(doc IDocument) map[string]interface{}

// IDocument is the common inherited interface
// for documents
type IDocument interface {
	GetTitle() string
	GetPath() string
	GetMime() string
	GetExtension() string
	GetMtime() time.Time
	SetTitle(t string)
	SetPath(p string)
	SetMime(m string)
	SetExtension(e string)
	SetMtime(t time.Time)
}

type Document struct {
	Title     string
	Path      string
	Mime      string
	Extension string
	Mtime     time.Time
}

func (doc Document) GetTitle() string {
	return doc.Title
}

func (doc Document) SetTitle(t string) {
	doc.Title = t
}

func (doc Document) GetPath() string {
	return doc.Path
}

func (doc Document) SetPath(p string) {
	doc.Path = p
}

func (doc Document) GetMime() string {
	return doc.Mime
}

func (doc Document) SetMime(m string) {
	doc.Mime = m
}

func (doc Document) GetExtension() string {
	return doc.Extension
}

func (doc Document) SetExtension(e string) {
	doc.Extension = e
}

func (doc Document) GetMtime() time.Time {
	return doc.Mtime
}

func (doc Document) SetMtime(t time.Time) {
	doc.Mtime = t
}
