package search

import "github.com/Zenika/RAO/document"

type SearchEngine interface {
	Store(index string, doc document.IDocument, docMapper document.DocumentMapper)
	Search(index string, query Query) (*Response, error)
	Configure(index string, settings map[string]interface{}) error
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

// type Settings struct {
// 	Index string
// 	Data  map[string]interface{}
// }
//
// type Options struct {
// 	Index string
// 	Data  map[string]interface{}
// }

type Response struct {
	Data interface{}
}

func (search SearchService) Store(index string, doc document.IDocument, docMapper document.DocumentMapper) {
	search.engine.Store(index, doc, docMapper)
}

// []byte
func (search SearchService) Search(index string, query Query) (*Response, error) {
	return search.engine.Search(index, query)
}

func (search SearchService) Configure(index string, settings map[string]interface{}) error {
	return search.engine.Configure(index, settings)
}
