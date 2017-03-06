package algolia

import (
	"encoding/json"
	"github.com/Zenika/RAO/auth"
	"github.com/Zenika/RAO/document"
	"github.com/Zenika/RAO/log"
	"github.com/Zenika/RAO/search"
	"github.com/algolia/algoliasearch-client-go/algoliasearch"
)

type Algolia struct {
	client algoliasearch.Client
}

func (alg Algolia) Store(documents []document.IDocument) {
	index := alg.client.InitIndex("rao")
	for _, doc := range documents {
		_, err := index.AddObject(
			algoliasearch.Object{
				"Content": doc.GetContent(),
				"Path":    doc.GetPath(),
				"Mime":    doc.GetMime(),
				"Mtime":   doc.GetMtime(),
				"Bytes":   doc.GetBytes(),
				"Client":  doc.GetClient(),
				"Region":  doc.GetRegion(),
				"Sum":     doc.GetSum(),
			})
		log.Error(err, log.ERROR)
	}

}

func (alg Algolia) Search(query search.SearchQuery) ([]byte, error) {
	index := alg.client.InitIndex("rao")
	settings := algoliasearch.Map{
		"facets":       query.Facets,
		"facetFilters": query.FacetFilters,
		"filters":      query.Filters,
		"page":         query.Page,
	}
	res, err := index.Search(query.Query, settings)
	recs, err := json.Marshal(res)
	if err == nil {
		return recs, nil
	} else {
		log.Error(err, log.ERROR)
		return nil, err
	}
}

func New() *Algolia {
	return &Algolia{
		client: auth.RequireAlgoliaClient(),
	}
}
