package algolia

import (
	"github.com/algolia/algoliasearch-client-go/algolia/search"
)

func NewSearchClient(appID, apiKey string) *search.Client {
	return NewSearchClientWithConfig(
		search.Configuration{
			AppID:  appID,
			ApiKey: apiKey,
		},
	)
}

func NewSearchClientWithConfig(config search.Configuration) *search.Client {
	return search.NewClient(config)
}
