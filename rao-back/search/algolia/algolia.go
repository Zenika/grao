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
	index  algoliasearch.Index
}

var index algoliasearch.Index = nil

func initIndex(client algoliasearch.Client, indexId string) algoliasearch.Index {
	index := client.InitIndex(indexId)
	settings := algoliasearch.Map{
		"attributesToRetrieve": []string{
			"Client",
			"Agence",
			"Content",
			"Mime",
			"Ext",
			"Bytes",
			"Sum",
			"Title",
			"Path",
		},
		"attributesForFaceting": []string{
			"Client",
			"Agence",
			"Ext",
		},
		"attributesToSnippet": []string{
			"Content:80",
		},
		"attributesToHighlight": []string{
			"Content",
		},
		"highlightPreTag":  `<em class="snippet">`,
		"highlightPostTag": "</em>",
	}
	_, err := index.SetSettings(settings)
	log.Error(err, log.ERROR)
	return index
}

func (alg Algolia) Store(documents []document.IDocument) {
	if nil == index {
		index = alg.client.InitIndex("rao")
	}
	for _, doc := range documents {
		_, err := index.AddObject(
			algoliasearch.Object{
				"Content": doc.GetContent(),
				"Path":    doc.GetPath(),
				"Mime":    doc.GetMime(),
				"Ext":     doc.GetExt(),
				"Mtime":   doc.GetMtime(),
				"Bytes":   doc.GetBytes(),
				"Client":  doc.GetClient(),
				"Agence":  doc.GetAgence(),
				"Sum":     doc.GetSum(),
			})
		log.Error(err, log.ERROR)
	}
}

func (alg Algolia) Search(query search.SearchQuery) ([]byte, error) {
	if nil == index {
		index = alg.client.InitIndex("rao")
	}
	settings := algoliasearch.Map{
		"facets":       query.Facets,
		"facetFilters": query.FacetFilters,
		"filters":      query.Filters,
		"page":         query.Page,
		// "typoTolerance": query.TypoTolerance,
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
	client := auth.RequireAlgoliaClient()
	index := initIndex(client, "rao")
	return &Algolia{
		client: client,
		index:  index,
	}
}
