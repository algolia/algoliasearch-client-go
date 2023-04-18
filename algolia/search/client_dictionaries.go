package search

import (
	"net/http"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/call"
)

// SaveDictionaryEntries saves dictionary entries for the given dictionary
func (c *Client) SaveDictionaryEntries(dictionaryName DictionaryName, dictionaryEntries []DictionaryEntry, opts ...interface{}) (res UpdateTaskRes, err error) {
	batchRequest := newAddDictionaryEntriesBatch(dictionaryEntries, false)
	path := c.path("/dictionaries/%s/batch", dictionaryName)
	err = c.transport.Request(&res, http.MethodPost, path, batchRequest, call.Write, opts...)
	res.wait = c.WaitTask
	return
}

// ReplaceDictionaryEntries replaces existing dictionary entries present in the dictionary with the given ones
func (c *Client) ReplaceDictionaryEntries(dictionaryName DictionaryName, dictionaryEntries []DictionaryEntry, opts ...interface{}) (res UpdateTaskRes, err error) {
	batchRequest := newAddDictionaryEntriesBatch(dictionaryEntries, true)
	path := c.path("/dictionaries/%s/batch", dictionaryName)
	err = c.transport.Request(&res, http.MethodPost, path, batchRequest, call.Write, opts...)
	res.wait = c.WaitTask
	return
}

// DeleteDictionaryEntries deletes dictionary entries with the given objectIDs
func (c *Client) DeleteDictionaryEntries(dictionaryName DictionaryName, objectIDs []string, opts ...interface{}) (res UpdateTaskRes, err error) {
	batchRequest := newDeleteDictionaryEntriesBatch(objectIDs, false)
	path := c.path("/dictionaries/%s/batch", dictionaryName)
	err = c.transport.Request(&res, http.MethodPost, path, batchRequest, call.Write, opts...)
	res.wait = c.WaitTask
	return
}

// ClearDictionaryEntries deletes all the dictionary entries from the given dictionary
func (c *Client) ClearDictionaryEntries(dictionaryName DictionaryName, opts ...interface{}) (res UpdateTaskRes, err error) {
	return c.ReplaceDictionaryEntries(dictionaryName, []DictionaryEntry{}, opts...)
}

// SearchDictionaryEntries searches for dictionary entries according to the given query string and any rule
func (c *Client) SearchDictionaryEntries(dictionaryName DictionaryName, query string, opts ...interface{}) (res SearchDictionariesRes, err error) {
	body := newSearchDictionariesParams(query, opts...)
	path := c.path("/dictionaries/%s/search", dictionaryName)
	err = c.transport.Request(&res, http.MethodPost, path, body, call.Read, opts...)
	return
}

// GetDictionarySettings Retrieve dictionaries settings
func (c *Client) GetDictionarySettings(opts ...interface{}) (res DictionarySettings, err error) {
	path := c.path("/dictionaries/*/settings")
	err = c.transport.Request(&res, http.MethodGet, path, nil, call.Read, opts...)
	return
}

// SetDictionarySettings Set dictionary settings
func (c *Client) SetDictionarySettings(settings DictionarySettings, opts ...interface{}) (res UpdateTaskRes, err error) {
	path := c.path("/dictionaries/*/settings")
	err = c.transport.Request(&res, http.MethodPut, path, settings, call.Write, opts...)
	res.wait = c.WaitTask
	return
}
