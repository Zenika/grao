package algolia

import (
	"fmt"

	"github.com/Zenika/RAO/auth"
	"github.com/Zenika/RAO/document"
	"github.com/Zenika/RAO/log"
	"github.com/Zenika/RAO/search"
	"github.com/algolia/algoliasearch-client-go/algoliasearch"
)

type Algolia struct {
	client algoliasearch.Client
	index  map[string]algoliasearch.Index
}

func New() *Algolia {
	client := auth.RequireAlgoliaClient()
	index := make(map[string]algoliasearch.Index)
	return &Algolia{
		client: client,
		index:  index,
	}
}

func (alg Algolia) Configure(index string, settings map[string]interface{}) error {
	i := alg.getIndex(index)
	_, err := i.SetSettings(settings)
	if nil != err {
		log.Error(err, log.ERROR)
	}
	return err

}

func (alg Algolia) Store(index string, doc document.IDocument, docMapper document.DocumentMapper) {
	i := alg.getIndex(index)
	alg.dedupe(i, doc)
	_, err := i.AddObject(docMapper(doc))
	log.Error(err, log.ERROR)
}

func (alg Algolia) Search(index string, query search.Query) (*search.Response, error) {
	i := alg.getIndex(index)
	if 0 == query.HitsPerPage {
		query.HitsPerPage = 20
	}
	settings := algoliasearch.Map{
		"facets":                       query.Facets,
		"facetFilters":                 query.FacetFilters,
		"filters":                      query.Filters,
		"page":                         query.Page,
		"hitsPerPage":                  query.HitsPerPage,
		"restrictSearchableAttributes": query.Restriction,
	}
	response, err := i.Search(query.Query, settings)
	if err == nil {
		return &(search.Response{Data: response}), err
	} else {
		log.Error(err, log.ERROR)
		return nil, err
	}
}

func (alg Algolia) getIndex(id string) algoliasearch.Index {
	if nil == alg.index[id] {
		alg.index[id] = alg.client.InitIndex(id)
	}
	return alg.index[id]
}

func (alg Algolia) dedupe(index algoliasearch.Index, doc document.IDocument) error {
	dups := fmt.Sprintf(`Path:"%s" AND Title:"%s"`, doc.GetPath(), doc.GetTitle())
	err := index.DeleteByQuery("", algoliasearch.Map{
		"filters": dups,
	})
	if nil != err {
		log.Error(err, log.ERROR)
	}
	return err
}
