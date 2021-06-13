package query_suggestions

import (
	"fmt"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/call"
	"net/http"
)

type IndexConfig struct {
	// Index name to target.
	IndexName string `json:"indexName"`
	// List of source indices used to generate a Query Suggestions index.
	SourceIndices []SourceIndex `json:"sourceIndices"`
	// De-duplicate singular and plural suggestions.
	Languages []string `json:"languages,omitempty"`
	// List of words and patterns to exclude from the Query Suggestions index.
	Exclude []string `json:"exclude,omitempty"`
}

type SourceIndex struct {
	// Source index name.
	IndexName string `json:"indexName"`
	// List of analytics tags to filter the popular searches per tag.
	AnalyticsTags []string `json:"analyticsTags,omitempty"`
	// List of facets to define as categories for the query suggestions
	Facets []map[string]interface{}
	// Minimum number of hits (e.g., matching records in the source index) to generate a suggestions.
	MinHits *int `json:"minHits"`
	// Minimum number of required letters for a suggestion to remain.
	MinLetters *int `json:"minLetters"`
	// List of facet attributes used to generate Query Suggestions. The resulting suggestions are every combination of the facets in the nested list
	Generate [][]string `json:"generate,omitempty"`
	// List of external indices to use to generate custom Query Suggestions.
	External []string `json:"external,omitempty"`
}

// CreateConfig creates new query suggestions index with given config.
func (c *Client) CreateConfig(config IndexConfig, opts ...interface{}) error {
	path := c.path("/configs")
	return c.transport.Request(map[string]interface{}{}, http.MethodPost, path, config, call.Write, opts...)
}

// UpdateConfig updates the query suggestions index config.
func (c *Client) UpdateConfig(config IndexConfig, opts ...interface{}) error {
	path := c.path(fmt.Sprintf("/configs/%s", config.IndexName))
	return c.transport.Request(map[string]interface{}{}, http.MethodPut, path, config, call.Write, opts...)
}

// DeleteConfig deletes the query suggestions index config.
func (c *Client) DeleteConfig(indexName string, opts ...interface{}) error {
	path := c.path(fmt.Sprintf("/configs/%s", indexName))
	return c.transport.Request(map[string]interface{}{}, http.MethodDelete, path, nil, call.Write, opts...)
}

// GetConfig retrieves the query suggestions index config by the given indexName.
func (c *Client) GetConfig(indexName string, opts ...interface{}) (config *IndexConfig, err error) {
	path := c.path(fmt.Sprintf("/configs/%s", indexName))
	err = c.transport.Request(&config, http.MethodGet, path, nil, call.Read, opts...)
	return
}

// ListConfigs lists all the configs of the query suggestions in a single call.
func (c *Client) ListConfigs(opts ...interface{}) (configs []*IndexConfig, err error) {
	path := c.path("/configs")
	err = c.transport.Request(&configs, http.MethodGet, path, nil, call.Read, opts...)
	return
}
