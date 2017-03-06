package search

import (
	"github.com/Zenika/RAO/document"
)

type SearchEngine interface {
	Store(documents []document.IDocument)
	Search(query SearchQuery) ([]byte, error)
}

type SearchService struct {
	engine SearchEngine
}

type SearchQuery struct {
	Query        string
	Facets       string
	FacetFilters string
	Filters      string
	Page         int
}

func New(eng SearchEngine) *SearchService {
	return &SearchService{
		engine: eng,
	}
}

func (search SearchService) Store(documents []document.IDocument) {
	search.engine.Store(documents)
}

func (search SearchService) Search(query SearchQuery) ([]byte, error) {
	return search.engine.Search(query)
}
