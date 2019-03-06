package search

import "errors"

var (
	NoMoreHitsErr error = errors.New("No more hits")
	// TODO: centralize most errors into this file
)
