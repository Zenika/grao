package search

import (
	"github.com/Zenika/RAO/document"
)

type SearchEngine interface {
	Store(documents []document.IDocument)
	Search(query Query) (*Response, error)
}

func New(eng SearchEngine) *SearchService {
	return &SearchService{
		engine: eng,
	}
}

type SearchService struct {
	engine SearchEngine
}

type Query struct {
	Query        string
	Facets       string
	FacetFilters string
	Filters      string
	Page         int
}

type Response struct {
	Data interface{}
}

func (search SearchService) Store(documents []document.IDocument) {
	search.engine.Store(documents)
}

// []byte
func (search SearchService) Search(query Query) (*Response, error) {
	return search.engine.Search(query)
}
