package search

import "errors"

var (
	NoMoreHitsErr         = errors.New("No more hits")
	NoMoreRulesErr        = errors.New("No more rules")
	NoMoreSynonymsErr     = errors.New("No more synonyms")
	SameAppIDErr          = errors.New("Indices cannot target the same application ID. Please use Client.CopyIndex for same-app index copy instead.")
	IndexAlreadyExistsErr = errors.New("Destination index already exists. Please delete it first as the CopyIndex cannot hold the responsibility of modifying the destination index.")
	// TODO: centralize most errors into this file
)
