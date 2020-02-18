package search

import (
	"net/http"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/call"
)

// SetPersonalizationStrategy defines and override the full configuration of
// the personalization strategy feature.
//
// Deprecated: use recommendation.Client.SetPersonalizationStrategy() instead
func (c *Client) SetPersonalizationStrategy(strategy Strategy, opts ...interface{}) (res SetPersonalizationStrategyRes, err error) {
	path := c.path("/recommendation/personalization/strategy")
	err = c.transport.Request(&res, http.MethodPost, path, strategy, call.Write, opts...)
	return
}

// GetPersonalizationStrategy retrieves the full configuration of the
// personalization strategy feature.
//
// Deprecated: use recommendation.Client.GetPersonalizationStrategy() instead
func (c *Client) GetPersonalizationStrategy(opts ...interface{}) (res GetPersonalizationStrategyRes, err error) {
	path := c.path("/recommendation/personalization/strategy")
	err = c.transport.Request(&res, http.MethodGet, path, nil, call.Read, opts...)
	return
}
