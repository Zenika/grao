package algolia

import(
  "github.com/algolia/algoliasearch-client-go/algoliasearch"
  "github.com/Zenika/RAO/dropbox"
  "github.com/Zenika/RAO/auth"
  "encoding/json"
  "log"
)

var alg algoliasearch.Client = auth.RequireAlgoliaClient()

func Push(documents []dropbox.DbxDocument){
  var objects []algoliasearch.Object = nil
  index := alg.InitIndex("rao")
  pld, err := json.Marshal(documents)
  err = json.Unmarshal(pld, &objects)
  check(err)
  _, err = index.AddObjects(objects)
  check(err)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
