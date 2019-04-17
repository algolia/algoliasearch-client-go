package search

import "errors"

var (
	NoMoreHitsErr         = errors.New("no more hits")
	NoMoreRulesErr        = errors.New("no more rules")
	NoMoreSynonymsErr     = errors.New("no more synonyms")
	SameAppIDErr          = errors.New("indices cannot target the same application ID, please use Client.CopyIndex for same-app index copy instead")
	IndexAlreadyExistsErr = errors.New("destination index already exists, please delete it first as the CopyIndex cannot hold the responsibility of modifying the destination index")
)
