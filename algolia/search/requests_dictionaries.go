package search

import (
	iopt "github.com/algolia/algoliasearch-client-go/v3/algolia/internal/opt"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/opt"
)

type dictionaryEntryRequest struct {
	Action string      `json:"action"`
	Body   interface{} `json:"body"`
}

type dictionaryEntryRequestsBatch struct {
	Requests                       []interface{} `json:"requests"`
	ClearExistingDictionaryEntries bool          `json:"clearExistingDictionaryEntries"`
}

func newAddDictionaryEntriesBatch(dictionaryEntries []DictionaryEntry, clearExistingDictionaryEntries bool) dictionaryEntryRequestsBatch {
	addRequests := make([]interface{}, len(dictionaryEntries))
	for i, entry := range dictionaryEntries {
		addRequests[i] = dictionaryEntryRequest{"addEntry", entry}
	}
	return dictionaryEntryRequestsBatch{addRequests, clearExistingDictionaryEntries}
}

func newDeleteDictionaryEntriesBatch(objectIDs []string, clearExistingDictionaryEntries bool) dictionaryEntryRequestsBatch {
	deleteRequests := make([]interface{}, len(objectIDs))
	for i, objectID := range objectIDs {
		deleteRequests[i] = dictionaryEntryRequest{
			"deleteEntry",
			struct {
				ObjectID string `json:"objectID"`
			}{
				objectID,
			},
		}
	}
	return dictionaryEntryRequestsBatch{deleteRequests, clearExistingDictionaryEntries}
}

type searchDictionariesParams struct {
	Query       string                 `json:"query"`
	Page        *opt.PageOption        `json:"page,omitempty"`
	HitsPerPage *opt.HitsPerPageOption `json:"hitsPerPage,omitempty"`
	Language    *opt.LanguageOption    `json:"language,omitempty"`
}

func newSearchDictionariesParams(query string, opts ...interface{}) searchDictionariesParams {
	return searchDictionariesParams{
		Query:       query,
		Page:        iopt.ExtractPage(opts...),
		HitsPerPage: iopt.ExtractHitsPerPage(opts...),
		Language:    iopt.ExtractLanguage(opts...),
	}
}
