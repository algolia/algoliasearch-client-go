package search

import (
	iopt "github.com/algolia/algoliasearch-client-go/algolia/internal/opt"
	"github.com/algolia/algoliasearch-client-go/algolia/opt"
)

type KeyQueryParams struct {
	RestrictSources *opt.RestrictSourcesOption `json:"restrictSources,omitempty"`
	QueryParams
}

func newKeyQueryParams(opts ...interface{}) KeyQueryParams {
	return KeyQueryParams{
		RestrictSources: iopt.ExtractRestrictSources(opts...),
		QueryParams:     newQueryParams(opts...),
	}
}
