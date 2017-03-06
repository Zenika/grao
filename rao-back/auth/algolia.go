package auth

import (
	"github.com/algolia/algoliasearch-client-go/algoliasearch"
	"os"
)

var alg algoliasearch.Client = nil

func RequireAlgoliaClient() algoliasearch.Client {
	if alg != nil {
		return alg
	}
	id := os.Getenv("RAO_ALGOLIA_ID")
	key := os.Getenv("RAO_ALGOLIA_KEY")
	alg := algoliasearch.NewClient(id, key)
	return alg
}
