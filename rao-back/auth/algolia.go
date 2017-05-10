package auth

import (
	"os"

	"github.com/algolia/algoliasearch-client-go/algoliasearch"
)

var alg algoliasearch.Client = nil

func RequireAlgoliaClient() algoliasearch.Client {
	if alg != nil {
		return alg
	}
	id := os.Getenv("GRAO_ALGOLIA_ID")
	key := os.Getenv("GRAO_ALGOLIA_KEY")
	alg := algoliasearch.NewClient(id, key)
	return alg
}
