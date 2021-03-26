package search

import "github.com/algolia/algoliasearch-client-go/v3/algolia/opt"

// DictionaryName represents a name of a dictionary containing linguistic resources provided by Algolia
type DictionaryName string

const (
	Stopwords DictionaryName = "stopwords"
	Compounds DictionaryName = "compounds"
	Plurals   DictionaryName = "plurals"
)

// DictionarySettings represents dictionaries settings.
type DictionarySettings struct {
	DisableStandardEntries *opt.DisableStandardEntriesOption `json:"disableStandardEntries,omitempty"`
}
