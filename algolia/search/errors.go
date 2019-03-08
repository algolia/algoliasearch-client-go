package search

import "errors"

var (
	NoMoreHitsErr  error = errors.New("No more hits")
	NoMoreRulesErr error = errors.New("No more rules")
	// TODO: centralize most errors into this file
)
