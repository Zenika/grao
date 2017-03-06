package search

import (
	"github.com/Zenika/RAO/dropbox"
)

type SearchEngine interface {
	Store(documents []dropbox.DbxDocument)
	Search(query SearchQuery) ([]byte, error)
}

type SearchService struct {
	engine SearchEngine
}

type SearchQuery struct {
	Query string
	Facets string
	FacetFilters string
	Filters string
	Page int
}

func (search SearchService) Store(documents []dropbox.DbxDocument) {
	search.engine.Store(documents)
}

func (search SearchService) Search(query SearchQuery) ([]byte, error) {
	return search.engine.Search(query)
}

func New(eng SearchEngine) *SearchService {
	return &SearchService{
		engine: eng,
	}
}
