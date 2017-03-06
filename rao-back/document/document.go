package document

import (
	"time"
)

type IDocument interface {
	GetTitle() string
	GetPath() string
	GetMime() string
	GetContent() string
	GetClient() string
	GetRegion() string
	GetMtime() time.Time
	GetBytes() int64
	GetSum() string
	SetTitle(t string)
	SetPath(p string)
	SetMime(m string)
	SetContent(c string)
	SetClient(c string)
	SetRegion(r string)
	SetMtime(t time.Time)
	SetBytes(b int64)
	SetSum(s string)
}

type Document struct {
	Title   string
	Path    string
	Mime    string
	Content string
	Client  string
	Region  string
	Mtime   time.Time
	Bytes   int64
	Sum     string
}

type DocumentHandler func(bytes []byte, doc IDocument)

// the title of the document (original name in FS hierarchy)
func (doc *Document) GetTitle() string {
	return doc.Title
}
func (doc *Document) GetPath() string {
	return doc.Path
}
func (doc *Document) GetMime() string {
	return doc.Mime
}
func (doc *Document) GetContent() string {
	return doc.Content
}
func (doc *Document) GetClient() string {
	return doc.Client
}
func (doc *Document) GetRegion() string {
	return doc.Region
}
func (doc *Document) GetMtime() time.Time {
	return doc.Mtime
}
func (doc *Document) GetSum() string {
	return doc.Sum
}
func (doc *Document) GetBytes() int64 {
	return doc.Bytes
}
func (doc *Document) SetTitle(t string) {
	doc.Title = t
}

func (doc *Document) SetPath(p string) {
	doc.Path = p
}
func (doc *Document) SetMime(m string) {
	doc.Mime = m
}
func (doc *Document) SetContent(c string) {
	doc.Content = c
}
func (doc *Document) SetClient(c string) {
	doc.Client = c
}
func (doc *Document) SetRegion(r string) {
	doc.Region = r
}
func (doc *Document) SetMtime(t time.Time) {
	doc.Mtime = t
}
func (doc *Document) SetSum(s string) {
	doc.Sum = s
}
func (doc *Document) SetBytes(b int64) {
	doc.Bytes = b
}
