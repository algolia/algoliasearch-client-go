package search

import (
	iopt "github.com/algolia/algoliasearch-client-go/v3/algolia/internal/opt"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/opt"
)

type searchSynonymsParams struct {
	Query       string                 `json:"query"`
	Page        *opt.PageOption        `json:"page,omitempty"`
	Type        *opt.TypeOption        `json:"type,omitempty"`
	HitsPerPage *opt.HitsPerPageOption `json:"hitsPerPage,omitempty"`
}

func newSearchSynonymsParams(query string, opts ...interface{}) searchSynonymsParams {
	return searchSynonymsParams{
		Query:       query,
		Page:        iopt.ExtractPage(opts...),
		Type:        iopt.ExtractType(opts...),
		HitsPerPage: iopt.ExtractHitsPerPage(opts...),
	}
}
