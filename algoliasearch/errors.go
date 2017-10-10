package algoliasearch

import "errors"

var (
	NoMoreHitsErr     error = errors.New("No more hits")
	NoMoreSynonymsErr error = errors.New("No more synonyms")
)
