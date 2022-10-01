package recommend

import (
	"net/http"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/call"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/search"
)

func (c *Client) GetRecommendations(options []RecommendationsOptions, opts ...interface{}) (res search.MultipleQueriesRes, err error) {
	path := c.path("")
	err = c.transport.Request(&res, http.MethodPost, path, multipleOptions{options}, call.Read)
	return
}

func (c *Client) GetRelatedProducts(options []RelatedProductsOptions, opts ...interface{}) (res search.MultipleQueriesRes, err error) {
	var requests []RecommendationsOptions
	for _, o := range options {
		requests = append(requests, o.recommendationsOptions)
	}
	return c.GetRecommendations(requests, opts)
}

func (c *Client) GetFrequentlyBoughtTogether(options []FrequentlyBoughtTogetherOptions, opts ...interface{}) (res search.MultipleQueriesRes, err error) {
	var requests []RecommendationsOptions
	for _, o := range options {
		requests = append(requests, o.recommendationsOptions)
	}
	return c.GetRecommendations(requests, opts)
}

func (c *Client) path(format string, a ...interface{}) string {
	return "/1/indexes/*/recommendations"
}
