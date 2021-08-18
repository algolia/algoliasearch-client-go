package dictionary

import (
	"net/http"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/call"
)

type EntryState string

const (
	EntryStateEnabled  EntryState = "enabled"
	EntryStateDisabled EntryState = "disabled"
)

type AddStopwordsEntry struct {
	ObjectID string     `json:"objectId"`
	Language string     `json:"language"`
	Word     string     `json:"word"`
	State    EntryState `json:"state"`
}

type AddPluralsEntry struct {
	ObjectID string     `json:"objectId"`
	Language string     `json:"language"`
	Words    string     `json:"words"`
	State    EntryState `json:"state"`
}

type AddCompoundEntry struct {
	ObjectID      string   `json:"objectId"`
	Language      string   `json:"language"`
	Word          string   `json:"word"`
	Decomposition []string `json:"decomposition"`
}

type DeleteEntry struct {
	ObjectID string `json:"objectId"`
}

type BatchStopwordsEntriesRequest struct {
	AddEntries    []*AddStopwordsEntry
	DeleteEntries []*DeleteEntry
	// When true, start the batch by removing all the custom entries from the dictionary.
	ClearExistingDictionaryEntries bool
}

type BatchPluralsEntriesRequest struct {
	AddEntries    []*AddPluralsEntry
	DeleteEntries []*DeleteEntry
	// When true, start the batch by removing all the custom entries from the dictionary.
	ClearExistingDictionaryEntries bool
}

type BatchCompoundEntriesRequest struct {
	AddEntries    []*AddCompoundEntry
	DeleteEntries []*DeleteEntry
	// When true, start the batch by removing all the custom entries from the dictionary.
	ClearExistingDictionaryEntries bool
}

type requestBody struct {
	Requests                       []*entry `json:"requests"`
	ClearExistingDictionaryEntries bool     `json:"clearExistingDictionaryEntries"`
}

type entry struct {
	Action string      `json:"action"`
	Body   interface{} `json:"body"`
}

// BatchStopwordsEntries sends a batch of stopwords entries.
func (c *Client) BatchStopwordsEntries(req BatchStopwordsEntriesRequest, opts ...interface{}) error {
	addEntries := make([]interface{}, 0, len(req.AddEntries))
	for _, addEntry := range req.AddEntries {
		addEntries = append(addEntries, addEntry)
	}
	reqBody := buildBatchEntryRequestBody(addEntries, req.DeleteEntries, req.ClearExistingDictionaryEntries)
	path := c.path("/dictionaries/stopwords/settings")
	return c.transport.Request(map[string]interface{}{}, http.MethodPost, path, reqBody, call.Write, opts...)
}

// BatchPluralsEntries sends a batch of plurals entries.
func (c *Client) BatchPluralsEntries(req BatchPluralsEntriesRequest, opts ...interface{}) error {
	addEntries := make([]interface{}, 0, len(req.AddEntries))
	for _, addEntry := range req.AddEntries {
		addEntries = append(addEntries, addEntry)
	}
	reqBody := buildBatchEntryRequestBody(addEntries, req.DeleteEntries, req.ClearExistingDictionaryEntries)
	path := c.path("/dictionaries/plurals/settings")
	return c.transport.Request(map[string]interface{}{}, http.MethodPost, path, reqBody, call.Write, opts...)
}

// BatchCompoundEntries sends a batch of compound entries.
func (c *Client) BatchCompoundEntries(req BatchCompoundEntriesRequest, opts ...interface{}) error {
	addEntries := make([]interface{}, 0, len(req.AddEntries))
	for _, addEntry := range req.AddEntries {
		addEntries = append(addEntries, addEntry)
	}
	reqBody := buildBatchEntryRequestBody(addEntries, req.DeleteEntries, req.ClearExistingDictionaryEntries)
	path := c.path("/dictionaries/compound/settings")
	return c.transport.Request(map[string]interface{}{}, http.MethodPost, path, reqBody, call.Write, opts...)
}

type StopwordsSettings struct {
	DisableStandardEntries map[string]bool `json:"disableStandardEntries"`
}

// GetSettings gets dictionary settings.
func (c *Client) GetSettings(opts ...interface{}) (settings *StopwordsSettings, err error) {
	path := c.path("/dictionaries/*/settings")
	err = c.transport.Request(&settings, http.MethodGet, path, settings, call.Write, opts...)
	return
}

// SetSettings sets dictionary settings.
func (c *Client) SetSettings(settings StopwordsSettings, opts ...interface{}) error {
	path := c.path("/dictionaries/*/settings")
	return c.transport.Request(map[string]interface{}{}, http.MethodPut, path, settings, call.Write, opts...)
}

func buildBatchEntryRequestBody(addEntries []interface{}, deleteEntries []*DeleteEntry, clearExistingDictionaryEntries bool) *requestBody {
	entries := make([]*entry, 0, len(addEntries)+len(deleteEntries))
	for _, addEntry := range addEntries {
		entries = append(entries, &entry{
			Action: "addEntry",
			Body:   addEntry,
		})
	}
	for _, deleteEntry := range deleteEntries {
		entries = append(entries, &entry{
			Action: "deleteEntry",
			Body:   deleteEntry,
		})
	}
	reqBody := requestBody{
		ClearExistingDictionaryEntries: clearExistingDictionaryEntries,
		Requests:                       entries,
	}
	return &reqBody
}
