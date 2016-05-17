package algoliasearch

// Client is a representation of an Algolia application. Once initialized it
// allows manipulations over the indexes of the application as well as
// network related parameters.
type Client interface {
	// SetExtraHeader allows to set custom headers while reaching out to
	// Algolia servers.
	SetExtraHeader(key, value string)

	// SetTimeout specifies timeouts to use with the HTTP connection.
	SetTimeout(connectTimeout, readTimeout int)

	// ListIndexes returns the list of all indexes belonging to this Algolia
	// application.
	ListIndexes() (indexes []IndexRes, err error)

	// InitIndex returns an Index object targeting `indexName`.
	InitIndex(name string) Index

	// ListKeys returns all the API keys available for this Algolia application.
	ListKeys() (keys []Key, err error)

	// MoveIndex renames the index named `source` as `destination`.
	MoveIndex(source, destination string) (UpdateTaskRes, error)

	// CopyIndex duplicates the index named `source` as `destination`.
	CopyIndex(source, destination string) (UpdateTaskRes, error)

	// AddKey creates a new API key from the supplied `ACL` and the specified
	// optional parameters.
	AddKey(ACL []string, params Map) (res AddKeyRes, err error)

	// UpdateKey updates the API key named `key` with the supplied
	// parameters.
	UpdateKey(key string, params Map) (res UpdateKeyRes, err error)

	// GetKey returns the ACL and validity of the API key named `key`.
	GetKey(key string) (res Key, err error)

	// DeleteKey deletes the API key named `key`.
	DeleteKey(key string) (res DeleteRes, err error)

	// GetLogs retrieves the `length` latest logs, starting at `offset`. Logs can
	// be filtered by type via `logType` being either "query", "build" or "error".
	GetLogs(params Map) (logs []LogRes, err error)

	// GenerateSecuredAPIKey generates a public API key intended to restrict access
	// to certain records.
	// This new key is built upon the existing key named `apiKey`. Tag filters
	// or query parameters used to restrict access to certain records are specified
	// via the `public` argument. A single `userToken` may be supplied, in order to
	// use rate limited access.
	GenerateSecuredAPIKey(apiKey string, params Map) (key string, err error)

	// MultipleQueries performs all the queries specified in `queries` and
	// aggregates the results. It accepts two additional arguments: the name of
	// the field used to store the index name in the queries, and the strategy used
	// to perform the multiple queries.
	// The strategy can either be "none" or "stopIfEnoughMatches".
	MultipleQueries(queries []IndexedQuery, strategy string) (res []MultipleQueryRes, err error)

	// Batch performs all queries in `queries`. Each query should contain the
	// targeted index, as well as the type of operation wanted.
	Batch(records []BatchOperationIndexed) (res MultipleBatchRes, err error)
}

// Index is the structure used to manipulate an Algolia index.
type Index interface {
	// Delete deletes the Algolia index.
	Delete() (res DeleteTaskRes, err error)

	// Clear removes every record from the Algolia index.
	Clear() (res UpdateTaskRes, err error)

	// GetObject retrieves the object as an interface representing the JSON-encoded
	// object. The `objectID` is used to uniquely identify the object in the index
	// while `attributes` contains the list of attributes to retrieve.
	GetObject(objectID string, attributes []string) (object Object, err error)

	// GetObjects retrieves the objects identified according to their `objectIDs`.
	GetObjects(objectIDs []string) (objects []Object, err error)

	// DeleteObject deletes an object from the index that is uniquely identified by
	// its `objectID`.
	DeleteObject(objectID string) (res DeleteTaskRes, err error)

	// GetSettings retrieves the index' settings.
	GetSettings() (settings Settings, err error)

	// SetSettings changes the index settings.
	SetSettings(settings Map) (res UpdateTaskRes, err error)

	// WaitTask stops the current execution until the task identified by its
	// `taskID` is finished. The waiting time between each check starts at 1s and
	// is increased by a factor of 2 at each retry (but is bounded at around
	// 20min).
	WaitTask(taskID int) error

	// ListKeys lists all the keys that can access the index.
	ListKeys() (keys []Key, err error)

	// AddKey creates a new API key from the supplied `ACL` and the specified
	// optional parameters for the current index.
	AddKey(ACL []string, params Map) (res AddKeyRes, err error)

	// UpdateKey updates the key identified by the value `k.Value` by replacing all
	// the original key's fields by the ones of `k`.
	UpdateKey(value string, k Key) (res UpdateKeyRes, err error)

	// GetKey retrieves the key identified by its `value` string.
	GetKey(value string) (key Key, err error)

	// DeleteKey deletes the key identified by its `value` string.
	DeleteKey(value string) (res DeleteRes, err error)

	// AddObject adds a new record to the index.
	AddObject(object Object) (res CreateObjectRes, err error)

	// UpdateObject replaces the record in the index matching the one given in
	// parameter, according to its `objectID`.
	UpdateObject(object Object) (res UpdateObjectRes, err error)

	// PartialUpdateObject modifies the record in the index matching the one given
	// in parameter, according to its `objectID`. However, the record is only
	// partially updated i.e. only the specified attributes will be
	// updated, the original record won't be replaced.
	PartialUpdateObject(object Object) (res UpdateTaskRes, err error)

	// AddObjects adds several objects to the index.
	AddObjects(objects []Object) (BatchRes, error)

	// UpdateObjects adds or replaces several objects at the same time, according
	// to their respective `objectID` attribute.
	UpdateObjects(objects []Object) (BatchRes, error)

	// PartialUpdateObjects partially updates several objects at the same time,
	// according to their respective `objectID` attribute.
	PartialUpdateObjects(objects []Object) (BatchRes, error)

	// DeleteObjects deletes several objects at the same time, according to their
	// respective `objectID` attribute.
	DeleteObjects(objectIDs []string) (BatchRes, error)

	// Batch processes all the specified `operations` in a batch manner. The
	// operations's actions could be one of the following:
	//   - addObject
	//   - updateObject
	//   - partialUpdateObject
	//   - partialUpdateObjectNoCreate
	//   - deleteObject
	//   - clear
	// For more informations, please refer to the official REST API documentation
	// available here: https://www.algolia.com/doc/rest#batch-write-operations.
	Batch(operations []BatchOperation) (res BatchRes, err error)

	// Copy copies the index into a new one called `name`.
	Copy(name string) (UpdateTaskRes, error)

	// Move renames the index into `name`.
	Move(name string) (UpdateTaskRes, error)

	// GetStatus returns the status of a task given its ID `taskID`.
	GetStatus(taskID int) (res TaskStatusRes, err error)

	// SearchSynonyms returns the synonyms matching `query` whose types match
	// `types`. To retrieve the first page, `page` should be set to 0.
	// `hitsPerPage` specifies how many synonym sets will be returned per page.
	SearchSynonyms(query string, types []string, page, hitsPerPage int) (synonyms []Synonym, err error)

	// GetSynonym retrieves the synonym identified by `objectID`.
	GetSynonym(objectID string) (s Synonym, err error)

	// AddSynonym adds the given `synonym` to be identified `objectID`. This
	// addition can be forwarded to the index slaves using `forwardToSlaves`.
	AddSynonym(objectID string, synonym Synonym, forwardToSlaves bool) (res UpdateTaskRes, err error)

	// DeleteSynonym removes the synonym identified by `objectID`. This deletion
	// can be forwarded to the index slaves of the index with `forwardToSlaves`.
	DeleteSynonym(objectID string, forwardToSlaves bool) (res DeleteTaskRes, err error)

	// ClearSynonyms removes all synonyms from the index. The clear operation can
	// be forwarded to the index slaves of the index using `forwardToSlaves`.
	ClearSynonyms(forwardToSlaves bool) (res UpdateTaskRes, err error)

	// BatchSynonyms adds all `synonyms` to the index. The index can be cleared
	// before by setting `replaceExistingSynonyms` to `true`. The optional clear
	// operation and the additions can be forwarded to the index slaves by setting
	// `forwardToSlaves` to `true'.
	BatchSynonyms(synonyms []Synonym, replaceExistingSynonyms, forwardToSlaves bool) (res UpdateTaskRes, err error)

	// Browse returns the hits found according to the given `params`. The result
	// also contains the cursor needed to paginate the result. This is a low-level
	// function, if you simply want to iterate through all the results, it is
	// preferable to use BrowseAll instead. For more informations about the Browse
	// endpoint, please refer to the REST API documentation:
	// https://www.algolia.com/doc/rest#browse-all-index-content
	Browse(params Map) (res BrowseRes, err error)

	// BrowseAll returns an iterator pointing to the first result that matches the
	// search query given the `params`. Calling `Next()` on the iterator will
	// returns all the hits one by one, without the 1000 elements limit of the
	// Search function.
	BrowseAll(params Map) (it IndexIterator, err error)

	// Search performs a search query according to the `query` search query and the
	// given `params`.
	Search(query string, params Map) (res QueryRes, err error)

	// DeleteByQuery finds all the records that match the `query`, according to the
	// given 'params` and deletes them.
	DeleteByQuery(query string, params Map) (res BatchRes, err error)
}

// IndexIterator is used by the BrowseAll functions to iterate over all the
// records of an index (or a subset according to what the query was).
type IndexIterator interface {
	// Next returns the next record each time is is called. Subsequent pages of
	// results are automatically loaded and an error is returned if a problem
	// arises. When the last element has been reached, an error is returned with
	// the following message: "No more hits".
	Next() (res Map, err error)
}
