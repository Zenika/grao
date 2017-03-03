package algolia

import(
  "github.com/algolia/algoliasearch-client-go/algoliasearch"
  "github.com/Zenika/RAO/dropbox"
  "github.com/Zenika/RAO/auth"
  "github.com/Zenika/RAO/log"
  "encoding/json"
)

type Algolia struct {
  client algoliasearch.Client
}

func (alg Algolia) Store(documents []dropbox.DbxDocument){
  var objects []algoliasearch.Object = nil
  index := alg.client.InitIndex("rao")
  // todo: replace
  pld, err := json.Marshal(documents)
  err = json.Unmarshal(pld, &objects)
  _, err = index.AddObjects(objects)
  log.Error(err, log.ERROR)
}

func (alg Algolia) Search(pattern string)([]byte, error){
  index := alg.client.InitIndex("rao")
  res, err := index.Search(pattern, nil)
  recs, err := json.Marshal(res)
  if err == nil {
      return recs, nil
  } else {
      log.Error(err, log.ERROR)
      return nil, err
  }
}



func New() *Algolia {
    return &Algolia {
        client: auth.RequireAlgoliaClient(),
    }
}
