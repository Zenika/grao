package search

import (
	"github.com/Zenika/RAO/document"
)

type SearchEngine interface {
	Store(documents []document.IDocument, options interface{})
	Search(query Query, options interface{}) (*Response, error)
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
	HitsPerPage  int
	Restriction  string
}

type Response struct {
	Data interface{}
}

func (search SearchService) Store(documents []document.IDocument, options interface{}) {
	search.engine.Store(documents, options)
}

// []byte
func (search SearchService) Search(query Query, options interface{}) (*Response, error) {
	return search.engine.Search(query, options)
}
