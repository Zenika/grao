package search

import (
	"github.com/Zenika/RAO/document"
)

type SearchEngine interface {
	Store(doc document.IDocument, docMapper document.DocumentMapper, options interface{})
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

func (search SearchService) Store(doc document.IDocument, docMapper document.DocumentMapper, options interface{}) {
	search.engine.Store(doc, docMapper, options)
}

// []byte
func (search SearchService) Search(query Query, options interface{}) (*Response, error) {
	return search.engine.Search(query, options)
}
