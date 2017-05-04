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

type SearchOptions struct {
	Index string
}

type IndexSettings struct {
	Index    string
	Settings algoliasearch.Map
}

func (alg Algolia) Configure(settings interface{}) error {
	index := alg.getIndex(settings.(IndexSettings).Index)
	_, err := index.SetSettings(settings.(IndexSettings).Settings)
	if nil != err {
		log.Error(err, log.ERROR)
	}
	return err

}

func (alg Algolia) getIndex(id string) algoliasearch.Index {
	if nil == alg.index[id] {
		alg.index[id] = alg.client.InitIndex(id)
	}
	return alg.index[id]
}

func (alg Algolia) dedupe(index algoliasearch.Index, documents []document.IDocument) error {
	for _, doc := range documents {
		dups := fmt.Sprintf(`Path:"%s" AND Title:"%s"`, doc.GetPath(), doc.GetTitle())
		err := index.DeleteByQuery("", algoliasearch.Map{
			"filters": dups,
		})
		if nil != err {
			log.Error(err, log.ERROR)
			return err
		}
	}
	return nil
}

func (alg Algolia) Store(documents []document.IDocument, options interface{}) {
	index := alg.getIndex(options.(SearchOptions).Index)
	alg.dedupe(index, documents)
	for _, doc := range documents {
		_, err := index.AddObject(
			algoliasearch.Object{
				"Content":   doc.GetContent(),
				"Title":     doc.GetTitle(),
				"Path":      doc.GetPath(),
				"Client":    doc.GetClient(),
				"Agence":    doc.GetAgence(),
				"Extension": doc.GetExtension(),
				"Mime":      doc.GetMime(),
				"Mtime":     doc.GetMtime(),
				"Bytes":     doc.GetBytes(),
				"Sum":       doc.GetSum(),
			})
		log.Error(err, log.ERROR)
	}
}

func (alg Algolia) Search(query search.Query, options interface{}) (*search.Response, error) {
	index := alg.getIndex(options.(SearchOptions).Index)
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
	response, err := index.Search(query.Query, settings)
	if err == nil {
		return &(search.Response{Data: response}), err
	} else {
		log.Error(err, log.ERROR)
		return nil, err
	}
}

func New() *Algolia {
	client := auth.RequireAlgoliaClient()
	// index := initIndex(client, "rao")
	index := make(map[string]algoliasearch.Index)
	return &Algolia{
		client: client,
		index:  index,
	}
}
