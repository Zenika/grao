package document

import "time"

type DocumentHandler func(doc IDocument) // TODO errror handling

type DocumentFilter func(doc IDocument) bool

type DocumentMapper func(doc IDocument) interface{}

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
