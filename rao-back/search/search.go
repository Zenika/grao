package search

import (
	"github.com/Zenika/RAO/dropbox"
)

type SearchEngine interface {
	Store(documents []dropbox.DbxDocument)
	Search(pattern string) ([]byte, error)
}

type SearchService struct {
	engine SearchEngine
}

func (search SearchService) Store(documents []dropbox.DbxDocument) {
	search.engine.Store(documents)
}

func (search SearchService) Search(pattern string) ([]byte, error) {
	return search.engine.Search(pattern)
}

func New(eng SearchEngine) *SearchService {
	return &SearchService{
		engine: eng,
	}
}
