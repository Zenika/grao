// algolia Package is a SearchEngine implementation
// used to store and search documents using the
// algolia cloud platform
//
// see https://www.algolia.com
package algolia

import (
	"fmt"

	"github.com/Zenika/rao/rao-back/auth"
	"github.com/Zenika/rao/rao-back/document"
	"github.com/Zenika/rao/rao-back/log"
	"github.com/Zenika/rao/rao-back/search"
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

func (alg Algolia) Configure(index string, settings search.Settings) error {
	i := alg.getIndex(index)
	_, err := i.SetSettings(algoliasearch.Map{
		"searchableAttributes":  settings.SearchableAttributes,
		"attributesToRetrieve":  settings.AttributesToRetrieve,
		"attributesForFaceting": settings.AttributesForFaceting,
		"attributesToSnippet":   settings.AttributesToSnippet,
		"attributesToHighlight": settings.AttributesToHighlight,
		"highlightPreTag":       settings.HighlightPreTag,
		"highlightPostTag":      settings.HighlightPostTag,
	})
	if nil != err {
		log.Debug("Y a une erreur la !")
		log.Error(err, log.ERROR)
	}
	log.Debug("On passe la !")
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
	log.Debug("[algolia.engine] - query : " + query.Restriction)
	settings := algoliasearch.Map{
		"facets":                       query.Facets,
		"facetFilters":                 query.FacetFilters,
		"filters":                      query.Filters,
		"page":                         query.Page,
		"hitsPerPage":                  query.HitsPerPage,
		"restrictSearchableAttributes": query.Restriction,
	}
	response, err := i.Search(query.Query, settings)

	fmt.Printf("response : %+v\n", response)

	fmt.Printf("err : %+v\n", err)

	if err == nil {
		return &(search.Response{Data: response}), err
	} else {
		fmt.Println("Y A UNE ERREUR !!!")
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
