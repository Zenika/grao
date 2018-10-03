// Search Package contains search service interfaces
// with subpackages related to their implementations.
//
// A search service compose a SearchEngine interface implementation,
// provided as an argument to the factory call
package search

import (
	"fmt"

	"github.com/Zenika/rao/rao-back/document"
)

// SearchEngine implementation own the responsability of
// implementing search service core methods
//
// **Store** stores the provided document to an index referenced
// by its first argument. docMapper function may be used to convert
// document to a map[string]interface{} complying with the underlying
// implementation
//
// **Search** will perform on an indexed referenced by its first argument
// a query provided as a seconde argument under the form of a Query object
//
// **Configure** should be used to tune index before performing queries if needed
type SearchEngine interface {
	Store(index string, doc document.IDocument, docMapper document.DocumentMapper)
	Search(index string, query Query) (*Response, error)
	Configure(index string, settings Settings) error
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

type Settings struct {
	SearchableAttributes  []string
	AttributesToRetrieve  []string
	AttributesForFaceting []string
	AttributesToSnippet   []string
	AttributesToHighlight []string
	HighlightPreTag       string
	HighlightPostTag      string
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
	fmt.Println("index :", index)
	fmt.Println("query :", query)
	return search.engine.Search(index, query)
}

func (search SearchService) Configure(index string, settings Settings) error {
	return search.engine.Configure(index, settings)
}
