package algolia

import(
  "github.com/algolia/algoliasearch-client-go/algoliasearch"
  "github.com/Zenika/RAO/dropbox"
  "github.com/Zenika/RAO/auth"
  "github.com/Zenika/RAO/log"
  "encoding/json"
)

var alg algoliasearch.Client = auth.RequireAlgoliaClient()

func Push(documents []dropbox.DbxDocument){
  var objects []algoliasearch.Object = nil
  index := alg.InitIndex("rao")
  // todo: replace
  pld, err := json.Marshal(documents)
  err = json.Unmarshal(pld, &objects)
  _, err = index.AddObjects(objects)
  log.Error(err, log.ERROR)
}
