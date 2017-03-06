package rao

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
  GetMtime() string
  GetBytes() string
  GetSum() string
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

// the title of the document (original name in FS hierarchy)
func (doc Document) GetTitle() (string){
  return doc.Title;
}

func (doc Document) GetPath() (string){
  return doc.Path;
}
func (doc Document) GetMime() (string){
  return doc.Mime;
}
func (doc Document) GetContent() (string){
  return doc.Content;
}
func (doc Document) GetClient() (string){
  return doc.Client;
}
func (doc Document) GetRegion() (string){
  return doc.Region;
}
func (doc Document) GetMtime() (time.Time){
  return doc.Mtime;
}
func (doc Document) GetSum() (string){
  return doc.Sum;
}
