package search

import "errors"

var (
	NoMoreHitsErr     = errors.New("No more hits")
	NoMoreRulesErr    = errors.New("No more rules")
	NoMoreSynonymsErr = errors.New("No more synonyms")
	// TODO: centralize most errors into this file
)
