package search

import "github.com/Zenika/RAO/document"

type SearchEngine interface {
	Store(index string, doc document.IDocument, docMapper document.DocumentMapper)
	Search(index string, query Query) (*Response, error)
	Configure(index string, settings map[string]interface{}) error
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

func New(eng SearchEngine) *SearchService {
	return &SearchService{
		engine: eng,
	}
}

func (search SearchService) Store(index string, doc document.IDocument, docMapper document.DocumentMapper) {
	search.engine.Store(index, doc, docMapper)
}

func (search SearchService) Search(index string, query Query) (*Response, error) {
	return search.engine.Search(index, query)
}

func (search SearchService) Configure(index string, settings map[string]interface{}) error {
	return search.engine.Configure(index, settings)
}
