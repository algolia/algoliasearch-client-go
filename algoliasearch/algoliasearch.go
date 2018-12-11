package algoliasearch

import (
	"net/http"
	"time"
)

// Client is a representation of an Algolia application. Once initialized it
// allows manipulations over the indexes of the application as well as network
// related parameters.
type Client interface {
	// GetAppID returns the Algolia application ID to which the client is linked to.
	GetAppID() string

	// SetExtraHeader allows to set custom headers while reaching out to
	// Algolia servers.
	SetExtraHeader(key, value string)

	// SetTimeout specifies timeouts to use with the HTTP connection.
	//
	// Deprecated: Use SetReadTimeout, SetWriteTimeout or SetAnalyticsTimeout
	// instead.
	SetTimeout(connectTimeout, readTimeout int)

	// SetReadTimeout allows to specify the timeout that will be used for all
	// read requests to the Algolia Search API.
	//
	// Default value is controlled by algoliasearch.DefaultReadTimeout.
	SetReadTimeout(t time.Duration)

	// SetWriteTimeout allows to specify the timeout that will be used for all
	// write (i.e. indexing) requests to the Algolia Search API.
	//
	// Default value is controlled by algoliasearch.DefaultWriteTimeout.
	SetWriteTimeout(t time.Duration)

	// SetAnalyticsTimeout allows to specify the timeout that will be used for
	// all analytics requests to the Algolia Analytics API.
	//
	// Default value is controlled by algoliasearch.DefaultAnalyticsTimeout.
	SetAnalyticsTimeout(t time.Duration)

	// SetMaxIdleConnsPerHosts specifies the value for `MaxIdleConnsPerHost` of
	// the underlying http.Transport.
	SetMaxIdleConnsPerHosts(maxIdleConnsPerHost int)

	// SetHTTPClient allows a custom HTTP client to be specified.
	// NOTE: using this may prevent timeouts set on this client from
	// working if the underlying transport is not of type *http.Transport.
	SetHTTPClient(client *http.Client)

	// ListIndexes returns the list of all indexes belonging to this Algolia
	// application.
	ListIndexes() (indexes []IndexRes, err error)

	// ListIndexesWithRequestOptions is the same as ListIndexes but it also
	// accepts extra RequestOptions.
	ListIndexesWithRequestOptions(opts *RequestOptions) (indexes []IndexRes, err error)

	// InitIndex returns an Index object targeting `name`.
	InitIndex(name string) Index

	// InitAnalytics returns a new Analytics instance, bound to the Client.
	InitAnalytics() Analytics

	// InitInsights returns a new Insights instance, bound to the Client.
	InitInsights() Insights

	// ListKeys returns all the API keys available for this Algolia
	// application.
	//
	// Deprecated: Use ListAPIKeys instead.
	ListKeys() (keys []Key, err error)

	// ListKeysWithRequestOptions is the same as ListKeys but it also accepts
	// extra RequestOptions.
	//
	// Deprecated: Use ListAPIKeysWithRequestOptions instead.
	ListKeysWithRequestOptions(opts *RequestOptions) (keys []Key, err error)

	// ListAPIKeys returns all the API keys available for this Algolia
	// application.
	ListAPIKeys() (keys []Key, err error)

	// ListAPIKeysWithRequestOptions is the same as ListKeys but it also
	// accepts extra RequestOptions.
	ListAPIKeysWithRequestOptions(opts *RequestOptions) (keys []Key, err error)

	// MoveIndex renames the index named `source` as `destination`.
	MoveIndex(source, destination string) (UpdateTaskRes, error)

	// MoveIndexWithRequestOptions is the same as MoveIndex but it also accepts
	// extra RequestOptions.
	MoveIndexWithRequestOptions(source, destination string, opts *RequestOptions) (UpdateTaskRes, error)

	// CopyIndex duplicates the index named `source` as `destination`.
	CopyIndex(source, destination string) (UpdateTaskRes, error)

	// CopyIndexWithRequestOptions is the same as CopyIndex but it also accepts
	// extra RequestOptions.
	CopyIndexWithRequestOptions(source, destination string, opts *RequestOptions) (UpdateTaskRes, error)

	// ScopedCopyIndex duplicates the index named `source` as `destination`,
	// according to the given scopes.
	ScopedCopyIndex(source, destination string, scopes []string) (UpdateTaskRes, error)

	// ScopedCopyIndexWithRequestOptions is the same as ScopedCopyIndex but it
	// also accepts extra RequestOptions.
	ScopedCopyIndexWithRequestOptions(source, destination string, scopes []string, opts *RequestOptions) (UpdateTaskRes, error)

	// DeleteIndex removes the `name` Algolia index.
	//
	// Deprecated: Use Index.Delete instead.
	DeleteIndex(name string) (res DeleteTaskRes, err error)

	// DeleteIndexWithRequestOptions is the same as DeleteIndex but it also
	// accepts extra RequestOptions.
	//
	// Deprecated: Use Index.DeleteWithRequestOptions instead.
	DeleteIndexWithRequestOptions(name string, opts *RequestOptions) (res DeleteTaskRes, err error)

	// ClearIndex removes every record from the `name` Algolia index.
	//
	// Deprecated: Use Index.Clear instead.
	ClearIndex(name string) (res UpdateTaskRes, err error)

	// ClearIndexWithRequestOptions is the same as ClearIndex but it also
	// accepts extra RequestOptions.
	//
	// Deprecated: Use Index.ClearWithRequestOptions instead.
	ClearIndexWithRequestOptions(name string, opts *RequestOptions) (res UpdateTaskRes, err error)

	// AddUserKey creates a new API key from the supplied `ACL` and the
	// specified optional parameters. More details here:
	// https://www.algolia.com/doc/rest#add-a-global-api-key
	//
	// Deprecated: Use AddAPiKey instead.
	AddUserKey(ACL []string, params Map) (AddKeyRes, error)

	// AddAPIKey creates a new API key from the supplied `ACL` and the
	// specified optional parameters. More details here:
	// https://www.algolia.com/doc/rest#add-a-global-api-key
	AddAPIKey(ACL []string, params Map) (res AddKeyRes, err error)

	// AddAPIKeyWithRequestOptions is the same as AddAPIKey but it also accepts
	// extra RequestOptions.
	AddAPIKeyWithRequestOptions(ACL []string, params Map, opts *RequestOptions) (res AddKeyRes, err error)

	// UpdateUserKey updates the API key identified by its value `key` with the
	// given parameters.
	//
	// Deprecated: Use UpdateAPIKey instead.
	UpdateUserKey(key string, params Map) (UpdateKeyRes, error)

	// UpdateAPIKey updates the API key identified by its value `key` with the
	// given parameters.
	UpdateAPIKey(key string, params Map) (res UpdateKeyRes, err error)

	// UpdateAPIKeyWithRequestOptions is the same as UpdateAPIKey but it also
	// accepts extra RequestOptions.
	UpdateAPIKeyWithRequestOptions(key string, params Map, opts *RequestOptions) (res UpdateKeyRes, err error)

	// GetUserKey returns the key identified by its value `key`.
	//
	// Deprecated: Use GetAPIKey instead.
	GetUserKey(key string) (Key, error)

	// GetAPIKey returns the key identified by its value `key`.
	GetAPIKey(key string) (res Key, err error)

	// GetAPIKeyWithRequestOptions is the same as GetAPIKey but it also accepts
	// extra RequestOptions.
	GetAPIKeyWithRequestOptions(key string, opts *RequestOptions) (res Key, err error)

	// DeleteUserKey deletes the API key identified by its `key`.
	//
	// Deprecated: Use DeleteAPIKey instead.
	DeleteUserKey(key string) (DeleteRes, error)

	// DeleteAPIKey deletes the API key identified by its `key`.
	DeleteAPIKey(key string) (res DeleteRes, err error)

	// DeleteAPIKeyWithRequestOptions is the same as DeleteAPIKey but it also
	// accepts extra RequestOptions.
	DeleteAPIKeyWithRequestOptions(key string, opts *RequestOptions) (res DeleteRes, err error)

	// GetLogs retrieves the logs according to the given `params` map which can
	// contain the following fields:
	//   - `length` (number of entries to retrieve)
	//   - `offset` (offset to the first entry)
	//   - `indexName` (index for which log entries should be retrieved)
	//   - `type` (type of logs to retrieve, can be "all", "query", "build" or
	//     "error")
	// More details here:
	// https://www.algolia.com/doc/rest-api/search/#get-last-logs
	GetLogs(params Map) (logs []LogRes, err error)

	// GetLogsWithRequestOptions is the same as GetLogs but it also accepts
	// extra RequestOptions.
	GetLogsWithRequestOptions(params Map, opts *RequestOptions) (logs []LogRes, err error)

	// MultipleQueries performs all the queries specified in `queries` and
	// aggregates the results. The `strategy` can either be set to `none`
	// (default) which executes all the queries until the last one, or set to
	// `stopIfEnoughMatches` to limit the number of results according to the
	// `hitsPerPage` parameter. More details here:
	// https://www.algolia.com/doc/rest#query-multiple-indexes
	MultipleQueries(queries []IndexedQuery, strategy string) (res []MultipleQueryRes, err error)

	// MultipleQueriesWithRequestOptions is the same as MultipleQueries but it
	// also accepts extra RequestOptions.
	MultipleQueriesWithRequestOptions(queries []IndexedQuery, strategy string, opts *RequestOptions) (res []MultipleQueryRes, err error)

	// List all available clusters for the application.
	ListClusters() (res []Cluster, err error)

	// ListClustersWithRequestOptions is the same as ListClusters but it
	// also accepts extra RequestOptions.
	ListClustersWithRequestOptions(opts *RequestOptions) (res []Cluster, err error)

	// List all user IDs across all clusters.
	ListUserIDs(page int, hitsPerPage int) (res ListUserIDsRes, err error)

	// ListUserIDsWithRequestOptions is the same as ListUserIDs but it
	// also accepts extra RequestOptions.
	ListUserIDsWithRequestOptions(page int, hitsPerPage int, opts *RequestOptions) (res ListUserIDsRes, err error)

	// Get a specific user ID when using multi cluster infrastructure.
	GetUserID(userID string) (res UserID, err error)

	// GetUserIDWithRequestOptions is the same as GetUserID but it
	// also accepts extra RequestOptions.
	GetUserIDWithRequestOptions(userID string, opts *RequestOptions) (res UserID, err error)

	// Assign a user ID to a cluster.
	AssignUserID(userID string, clusterName string) (res AssignUserIDRes, err error)

	// AssignUserIDWithRequestOptions is the same as AssignUserID but it
	// also accepts extra RequestOptions.
	AssignUserIDWithRequestOptions(userID string, clusterName string, opts *RequestOptions) (res AssignUserIDRes, err error)

	// Remove a user ID from a cluster.
	RemoveUserID(userID string) (res RemoveUserIDRes, err error)

	// RemoveUserIDWithRequestOptions is the same as RemoveUserID but it
	// also accepts extra RequestOptions.
	RemoveUserIDWithRequestOptions(userID string, opts *RequestOptions) (res RemoveUserIDRes, err error)

	// Get user IDs with the highest number of records per cluster.
	GetTopUserIDs() (res TopUserIDs, err error)

	// GetTopUserIDsWithRequestOptions is the same as GetTopUserIDs but it
	// also accepts extra RequestOptions.
	GetTopUserIDsWithRequestOptions(opts *RequestOptions) (res TopUserIDs, err error)

	// Typical Algolia search but for user ID.
	SearchUserIDs(query string, params Map) (res SearchUserIDRes, err error)

	// SearchUserIDsWithRequestOptions is the same as SearchUserIDs but it
	// also accepts extra RequestOptions.
	SearchUserIDsWithRequestOptions(query string, params Map, opts *RequestOptions) (res SearchUserIDRes, err error)

	// Batch performs all queries in `operations`.
	Batch(operations []BatchOperationIndexed) (res MultipleBatchRes, err error)

	// BatchWithRequestOptions is the same as Batch but it also accepts extra
	// RequestOptions.
	BatchWithRequestOptions(operations []BatchOperationIndexed, opts *RequestOptions) (res MultipleBatchRes, err error)

	// WaitTask stops the current execution until the task identified by its
	// `taskID` on the index `indexName` is finished. The waiting time between each check is usually
	// implemented by starting at 1s and increases by a factor of 2 at each
	// retry (but is bounded at around 20min).
	WaitTask(indexName string, taskID int) error

	// WaitTaskWithRequestOptions is the same as WaitTask but it also accepts
	// extra RequestOptions.
	WaitTaskWithRequestOptions(indexName string, taskID int, opts *RequestOptions) error

	// GetStatus returns the status of a task given its ID `taskID` and `indexName`.
	GetStatus(indexName string, taskID int) (res TaskStatusRes, err error)

	// GetStatusWithRequestOptions is the same as GetStatus but it also accepts
	// extra RequestOptions.
	GetStatusWithRequestOptions(indexName string, taskID int, opts *RequestOptions) (res TaskStatusRes, err error)

	// CopySettings copies the settings from the source index to the destination index.
	CopySettings(source, destination string) (UpdateTaskRes, error)

	// CopySettingsWithRequestOptions is the same as CopySettings but it also accepts
	// extra RequestOptions.
	CopySettingsWithRequestOptions(source, destination string, opts *RequestOptions) (UpdateTaskRes, error)

	// CopySynonyms copies the synonyms from the source index to the destination index.
	CopySynonyms(source, destination string) (UpdateTaskRes, error)

	// CopySynonymsWithRequestOptions is the same as CopySynonyms but it also accepts
	// extra RequestOptions.
	CopySynonymsWithRequestOptions(source, destination string, opts *RequestOptions) (UpdateTaskRes, error)

	// CopyRules copies the rules from the source index to the destination index.
	CopyRules(source, destination string) (UpdateTaskRes, error)

	// CopyRulesWithRequestOptions is the same as CopyRulesWith but it also accepts
	// extra RequestOptions.
	CopyRulesWithRequestOptions(source, destination string, opts *RequestOptions) (UpdateTaskRes, error)

	// SetPersonalizationStrategy allows to set the inner strategies related to the
	// Personalization features.
	SetPersonalizationStrategy(strategy Strategy) (SetStrategyRes, error)

	// SetPersonalizationStrategyWithRequestOptions is the same as SetPersonalizationStrategy but it also
	// accepts extra RequestOptions.
	SetPersonalizationStrategyWithRequestOptions(strategy Strategy, opts *RequestOptions) (SetStrategyRes, error)

	// GetPersonalizationStrategy retrieves the inner strategies related to the Personalization
	// features.
	GetPersonalizationStrategy() (Strategy, error)

	// GetPersonalizationStrategyWithRequestOptions is the same as GetPersonalizationStrategy but it also
	// accepts extra RequestOptions.
	GetPersonalizationStrategyWithRequestOptions(opts *RequestOptions) (Strategy, error)
}

// Index is a representation used to manipulate an Algolia index.
type Index interface {
	// GetAppID returns the Algolia application ID to which the index is linked to.
	GetAppID() string

	// Delete removes the Algolia index.
	Delete() (res DeleteTaskRes, err error)

	// DeleteWithRequestOptions is the same as Delete but it also accepts extra
	// RequestOptions.
	DeleteWithRequestOptions(opts *RequestOptions) (res DeleteTaskRes, err error)

	// Clear removes every record from the index.
	Clear() (res UpdateTaskRes, err error)

	// ClearWithRequestOptions is the same as Clear but it also accepts extra
	// RequestOptions.
	ClearWithRequestOptions(opts *RequestOptions) (res UpdateTaskRes, err error)

	// GetObject retrieves the object as an interface representing the
	// JSON-encoded object. The `objectID` is used to uniquely identify the
	// object in the index while `attributes` contains the list of attributes
	// to retrieve.
	GetObject(objectID string, attributes []string) (object Object, err error)

	// GetObjectWithRequestOptions is the same as GetObject but it also accepts
	// extra RequestOptions.
	GetObjectWithRequestOptions(objectID string, attributes []string, opts *RequestOptions) (object Object, err error)

	// GetObjects retrieves the objects identified according to their
	// `objectIDs`.
	GetObjects(objectIDs []string) (objects []Object, err error)

	// GetObjectsWithRequestOptions is the same as GetObjects but it also
	// accepts extra RequestOptions.
	GetObjectsWithRequestOptions(objectIDs []string, opts *RequestOptions) (objects []Object, err error)

	// GetObjectsAttrs retrieves the selected attributes of the objects
	// identified according to their `objectIDs`.
	GetObjectsAttrs(objectIDs, attributesToRetrieve []string) (objs []Object, err error)

	// GetObjectsAttrsWithRequestOptions is the same as GetObjectsAttrs but it
	// also accepts extra RequestOptions.
	GetObjectsAttrsWithRequestOptions(objectIDs, attributesToRetrieve []string, opts *RequestOptions) (objs []Object, err error)

	// DeleteObject deletes an object from the index that is uniquely
	// identified by its `objectID`.
	DeleteObject(objectID string) (res DeleteTaskRes, err error)

	// DeleteObjectWithRequestOptions is the same as DeleteObject but it also
	// accepts extra RequestOptions.
	DeleteObjectWithRequestOptions(objectID string, opts *RequestOptions) (res DeleteTaskRes, err error)

	// GetSettings retrieves the index's settings.
	GetSettings() (settings Settings, err error)

	// GetSettingsWithRequestOptions is the same as GetSettings but it also
	// accepts extra RequestOptions.
	GetSettingsWithRequestOptions(opts *RequestOptions) (settings Settings, err error)

	// SetSettings changes the index settings.
	SetSettings(settings Map) (res UpdateTaskRes, err error)

	// SetSettingsWithRequestOptions is the same as SetSettings but it also
	// accepts extra RequestOptions.
	SetSettingsWithRequestOptions(settings Map, opts *RequestOptions) (res UpdateTaskRes, err error)

	// WaitTask stops the current execution until the task identified by its
	// `taskID` is finished. The waiting time between each check is usually
	// implemented by starting at 1s and increases by a factor of 2 at each
	// retry (but is bounded at around 20min).
	WaitTask(taskID int) error

	// WaitTaskWithRequestOptions is the same as WaitTask but it also accepts
	// extra RequestOptions.
	WaitTaskWithRequestOptions(taskID int, opts *RequestOptions) error

	// ListKeys lists all the keys that can access the index.
	//
	// Deprecated: Use Client.ListAPIKeys instead.
	ListKeys() (keys []Key, err error)

	// ListKeysWithRequestOptions is the same as ListKeys but it also accepts
	// extra RequestOptions.
	//
	// Deprecated: Use Client.ListAPIKeysWithRequestOptions instead.
	ListKeysWithRequestOptions(opts *RequestOptions) (keys []Key, err error)

	// AddUserKey creates a new API key from the supplied `ACL` and the
	// specified optional `params` parameters for the current index. More
	// details here:
	// https://www.algolia.com/doc/rest#add-an-index-specific-api-key
	//
	// Deprecated: Use Client.AddAPIKey instead.
	AddUserKey(ACL []string, params Map) (AddKeyRes, error)

	// AddAPIKey creates a new API key from the supplied `ACL` and the
	// specified optional `params` parameters for the current index. More
	// details here:
	// https://www.algolia.com/doc/rest#add-an-index-specific-api-key
	//
	// Deprecated: Use Client.AddAPIKey instead.
	AddAPIKey(ACL []string, params Map) (res AddKeyRes, err error)

	// AddAPIKeyWithRequestOptions is the same as AddAPIKey but it also accepts
	// extra RequestOptions.
	//
	// Deprecated: Use Client.AddAPIKeyWithRequestOptions instead.
	AddAPIKeyWithRequestOptions(ACL []string, params Map, opts *RequestOptions) (res AddKeyRes, err error)

	// UpdateUserKey updates the key identified by its `key` with all the fields
	// present in the `params` Map. More details here:
	// https://www.algolia.com/doc/rest#update-an-index-specific-api-key
	//
	// Deprecated: Use Client.UpdateAPIKey instead.
	UpdateUserKey(key string, params Map) (UpdateKeyRes, error)

	// UpdateAPIKey updates the key identified by its `key` with all the fields
	// present in the `params` Map. More details here:
	// https://www.algolia.com/doc/rest#update-an-index-specific-api-key
	//
	// Deprecated: Use Client.UpdateAPIKey instead.
	UpdateAPIKey(key string, params Map) (res UpdateKeyRes, err error)

	// UpdateAPIKeyWithRequestOptions is the same as UpdateAPIKey but it also
	// accepts extra RequestOptions.
	//
	// Deprecated: Use Client.UpdateAPIKeyWithRequestOptions instead.
	UpdateAPIKeyWithRequestOptions(key string, params Map, opts *RequestOptions) (res UpdateKeyRes, err error)

	// GetUserKey retrieves the key identified by its `value`.
	//
	// Deprecated: Use Client.GetAPIKey instead.
	GetUserKey(value string) (Key, error)

	// GetAPIKey retrieves the key identified by its `value`.
	//
	// Deprecated: Use Client.GetAPIKey instead.
	GetAPIKey(value string) (key Key, err error)

	// GetAPIKeyWithRequestOptions is the same as GetAPIKey but it also accepts
	// extra RequestOptions.
	//
	// Deprecated: Use Client.GetAPIKeyWithRequestOptions instead.
	GetAPIKeyWithRequestOptions(value string, opts *RequestOptions) (key Key, err error)

	// DeleteUserKey deletes the key identified by its `value`.
	//
	// Deprecated: Use Client.DeleteAPIKey instead.
	DeleteUserKey(value string) (DeleteRes, error)

	// DeleteAPIKey deletes the key identified by its `value`.
	//
	// Deprecated: Use Client.DeleteAPIKey instead.
	DeleteAPIKey(value string) (res DeleteRes, err error)

	// DeleteAPIKeyWithRequestOptions is the same as DeleteAPIKey but it also
	// accepts extra RequestOptions.
	//
	// Deprecated: Use Client.DeleteAPIKeyWithRequestOptions instead.
	DeleteAPIKeyWithRequestOptions(value string, opts *RequestOptions) (res DeleteRes, err error)

	// AddObject adds a new record to the index.
	AddObject(object Object) (res CreateObjectRes, err error)

	// AddObjectWithRequestOptions is the same as AddObject but it also accepts
	// extra RequestOptions.
	AddObjectWithRequestOptions(object Object, opts *RequestOptions) (res CreateObjectRes, err error)

	// UpdateObject replaces the record in the index matching the one given in
	// parameter, according to its `objectID` attribute.
	UpdateObject(object Object) (res UpdateObjectRes, err error)

	// UpdateObjectWithRequestOptions is the same as UpdateObject but it also
	// accepts extra RequestOptions.
	UpdateObjectWithRequestOptions(object Object, opts *RequestOptions) (res UpdateObjectRes, err error)

	// PartialUpdateObject modifies the record in the index matching the one
	// given in parameter, according to its `objectID` attribute. However, the
	// record is only partially updated i.e. only the specified attributes will
	// be updated, the original record won't be replaced.
	PartialUpdateObject(object Object) (res UpdateTaskRes, err error)

	// PartialUpdateObjectWithRequestOptions is the same as PartialUpdateObject
	// but it also accepts extra RequestOptions.
	PartialUpdateObjectWithRequestOptions(object Object, opts *RequestOptions) (res UpdateTaskRes, err error)

	// PartialUpdateObjectNoCreate modifies the record in the index matching
	// the one given in parameter, according to its `objectID` attribute with a
	// partial update. However, if the object does not exist in the Algolia
	// index, the object is not created.
	PartialUpdateObjectNoCreate(object Object) (res UpdateTaskRes, err error)

	// PartialUpdateObjectNoCreateWithRequestOptions is the same as
	// PartialUpdateObjectNoCreate but it also accepts extra RequestOptions.
	PartialUpdateObjectNoCreateWithRequestOptions(object Object, opts *RequestOptions) (res UpdateTaskRes, err error)

	// AddObjects adds several objects to the index.
	AddObjects(objects []Object) (BatchRes, error)

	// AddObjectsWithRequestOptions is the same as AddObjects but it also
	// accepts extra RequestOptions.
	AddObjectsWithRequestOptions(objects []Object, opts *RequestOptions) (BatchRes, error)

	// UpdateObjects adds or replaces several objects at the same time,
	// according to their respective `objectID` attribute.
	UpdateObjects(objects []Object) (BatchRes, error)

	// UpdateObjectsWithRequestOptions is the same as UpdateObjects but it also
	// accepts extra RequestOptions.
	UpdateObjectsWithRequestOptions(objects []Object, opts *RequestOptions) (BatchRes, error)

	// PartialUpdateObjects partially updates several objects at the same time,
	// according to their respective `objectID` attribute.
	PartialUpdateObjects(objects []Object) (BatchRes, error)

	// PartialUpdateObjectsWithRequestOptions is the same as
	// PartialUpdateObjects but it also accepts extra RequestOptions.
	PartialUpdateObjectsWithRequestOptions(objects []Object, opts *RequestOptions) (BatchRes, error)

	// PartialUpdateObjectsNoCreate partially updates several objects at the
	// same time, according to their respective `objectID` attribute, but does
	// not create them if they do not exist.
	PartialUpdateObjectsNoCreate(objects []Object) (BatchRes, error)

	// PartialUpdateObjectsNoCreateWithRequestOptions is the same as
	// PartialUpdateObjectsNoCreate but it also accepts extra RequestOptions.
	PartialUpdateObjectsNoCreateWithRequestOptions(objects []Object, opts *RequestOptions) (BatchRes, error)

	// DeleteObjects removes several objects at the same time, according to
	// their respective `objectID` attribute.
	DeleteObjects(objectIDs []string) (BatchRes, error)

	// DeleteObjectsWithRequestOptions is the same as DeleteObjects but it also
	// accepts extra RequestOptions.
	DeleteObjectsWithRequestOptions(objectIDs []string, opts *RequestOptions) (BatchRes, error)

	// Batch processes all the specified `operations` in a batch manner. The
	// operations's actions could be one of the following:
	//   - `addObject`
	//   - `updateObject`
	//   - `partialUpdateObject`
	//   - `partialUpdateObjectNoCreate`
	//   - `deleteObject`
	//   - `clear`
	// More details here:
	// https://www.algolia.com/doc/rest#batch-write-operations.
	Batch(operations []BatchOperation) (res BatchRes, err error)

	// BatchWithRequestOptions is the same as Batch but it also accepts extra
	// RequestOptions.
	BatchWithRequestOptions(operations []BatchOperation, opts *RequestOptions) (res BatchRes, err error)

	// Copy copies the index into a new one called `name`.
	//
	// Deprecated: Use Client.CopyIndex instead.
	Copy(name string) (UpdateTaskRes, error)

	// CopyWithRequestOptions is the same as Copy but it also accepts extra
	// RequestOptions.
	//
	// Deprecated: Use Client.CopyIndexWithRequestOptions instead.
	CopyWithRequestOptions(name string, opts *RequestOptions) (UpdateTaskRes, error)

	// ScopedCopy copies the index into a new one called `name`, according to
	// the given scopes.
	//
	// Deprecated: Use Client.ScopedCopyIndex instead.
	ScopedCopy(name string, scopes []string) (UpdateTaskRes, error)

	// ScopedCopyWithRequestOptions is the same as ScopedCopy but it also
	// accepts extra RequestOptions.
	//
	// Deprecated: Use Client.ScopedCopyIndexWithRequestOptions instead.
	ScopedCopyWithRequestOptions(name string, scopes []string, opts *RequestOptions) (UpdateTaskRes, error)

	// Move renames the index into `name`.
	//
	// Deprecated: Use Client.MoveIndex instead.
	Move(name string) (UpdateTaskRes, error)

	// MoveWithRequestOptions is the same as Move but it also accepts extra
	// RequestOptions.
	//
	// Deprecated: Use Client.MoveIndexWithRequestOptions instead.
	MoveWithRequestOptions(name string, opts *RequestOptions) (UpdateTaskRes, error)

	// MoveTo renames the index into `name`.
	//
	// Deprecated: Use Client.MoveIndex instead.
	MoveTo(name string) (UpdateTaskRes, error)

	// MoveToWithRequestOptions is the same as MoveTo but it also accepts extra
	// RequestOptions.
	//
	// Deprecated: Use Client.MoveIndexWithRequestOptions instead.
	MoveToWithRequestOptions(name string, opts *RequestOptions) (UpdateTaskRes, error)

	// GetStatus returns the status of a task given its ID `taskID`.
	GetStatus(taskID int) (res TaskStatusRes, err error)

	// GetStatusWithRequestOptions is the same as GetStatus but it also accepts
	// extra RequestOptions.
	GetStatusWithRequestOptions(taskID int, opts *RequestOptions) (res TaskStatusRes, err error)

	// SearchSynonyms returns the synonyms matching `query` whose types match
	// `types`. To retrieve the first page, `page` should be set to 0.
	// `hitsPerPage` specifies how many synonym sets will be returned per page.
	SearchSynonyms(query string, types []string, page, hitsPerPage int) (synonyms []Synonym, err error)

	// SearchSynonymsWithRequestOptions is the same as SearchSynonyms but it
	// also accepts extra RequestOptions.
	SearchSynonymsWithRequestOptions(query string, types []string, page, hitsPerPage int, opts *RequestOptions) (synonyms []Synonym, err error)

	// GetSynonym retrieves the synonym identified by its `objectID`.
	GetSynonym(objectID string) (s Synonym, err error)

	// GetSynonymWithRequestOptions is the same as GetSynonym but it also
	// accepts extra RequestOptions.
	GetSynonymWithRequestOptions(objectID string, opts *RequestOptions) (s Synonym, err error)

	// AddSynonym adds the given `synonym`. This addition can be forwarded to
	// the index replicas by setting `forwardToReplicas` to `true`.
	//
	// Deprecated: Use SaveSynonym instead.
	AddSynonym(synonym Synonym, forwardToReplicas bool) (res UpdateTaskWithIDRes, err error)

	// AddSynonymWithRequestOptions is the same as AddSynonym but it also
	// accepts extra RequestOptions.
	//
	// Deprecated: Use SaveSynonymWithRequestOptions instead.
	AddSynonymWithRequestOptions(synonym Synonym, forwardToReplicas bool, opts *RequestOptions) (res UpdateTaskWithIDRes, err error)

	// SaveSynonym adds the given `synonym`. This addition can be forwarded to
	// the index replicas by setting `forwardToReplicas` to `true`.
	SaveSynonym(synonym Synonym, forwardToReplicas bool) (res UpdateTaskWithIDRes, err error)

	// SaveSynonymWithRequestOptions is the same as SaveSynonym but it also
	// accepts extra RequestOptions.
	SaveSynonymWithRequestOptions(synonym Synonym, forwardToReplicas bool, opts *RequestOptions) (res UpdateTaskWithIDRes, err error)

	// DeleteSynonym removes the synonym identified by its `objectID`. This
	// deletion can be forwarded to the index replicas by setting
	// `forwardToReplicas` to `true`.
	DeleteSynonym(objectID string, forwardToReplicas bool) (res DeleteTaskRes, err error)

	// DeleteSynonymWithRequestOptions is the same as DeleteSynonym but it also
	// accepts extra RequestOptions.
	DeleteSynonymWithRequestOptions(objectID string, forwardToReplicas bool, opts *RequestOptions) (res DeleteTaskRes, err error)

	// ClearSynonyms removes all synonyms from the index. The clear operation
	// can be forwarded to the index replicas by setting `forwardToReplicas` to
	// `true`.
	ClearSynonyms(forwardToReplicas bool) (res UpdateTaskRes, err error)

	// ClearSynonymsWithRequestOptions is the same as ClearSynonyms but it also
	// accepts extra RequestOptions.
	ClearSynonymsWithRequestOptions(forwardToReplicas bool, opts *RequestOptions) (res UpdateTaskRes, err error)

	// BatchSynonyms adds all `synonyms` to the index. The index can be cleared
	// before by setting `replaceExistingSynonyms` to `true`. The optional
	// clear operation and the additions can be forwarded to the index replicas
	// by setting `forwardToReplicas` to `true'.
	BatchSynonyms(synonyms []Synonym, replaceExistingSynonyms, forwardToReplicas bool) (res UpdateTaskRes, err error)

	// BatchSynonymsWithRequestOptions is the same as BatchSynonyms but it also
	// accepts extra RequestOptions.
	BatchSynonymsWithRequestOptions(synonyms []Synonym, replaceExistingSynonyms, forwardToReplicas bool, opts *RequestOptions) (res UpdateTaskRes, err error)

	// Browse returns the hits found according to the given `params`. The
	// `cursor` parameter controls the pagination of the results that `Browse`
	// is able to load. The first time `Browse` is called, `cursor` should be
	// an empty string. After that, subsequent calls must used the updated
	// `cursor` received in the response. This is a low-level function, if you
	// simply want to iterate through all the results, it is preferable to use
	// `BrowseAll` instead. More details here:
	// https://www.algolia.com/doc/rest#browse-all-index-content
	Browse(params Map, cursor string) (res BrowseRes, err error)

	// BrowseWithRequestOptions is the same as Browse but it also accepts extra
	// RequestOptions.
	BrowseWithRequestOptions(params Map, cursor string, opts *RequestOptions) (res BrowseRes, err error)

	// BrowseAll returns an iterator pointing to the first result that matches
	// the search query given the `params`. Calling `Next()` on the iterator
	// will returns all the hits one by one, without the 1000 elements limit of
	// the Search function. Once the last element as been reached, the next
	// call to `Next()` will return a `NoMoreHitsErr` error. If anything went
	// wrong during the browsing, a non-nil error is also returned.
	BrowseAll(params Map) (it IndexIterator, err error)

	// BrowseAllWithRequestOptions is the same as BrowseAll but it also accepts
	// extra RequestOptions.
	BrowseAllWithRequestOptions(params Map, opts *RequestOptions) (it IndexIterator, err error)

	// Search performs a search query according to the `query` search query and
	// the given `params`. More details here:
	// https://www.algolia.com/doc/rest#query-an-index
	Search(query string, params Map) (res QueryRes, err error)

	// SearchWithRequestOptions is the same as Search but it also accepts extra
	// RequestOptions.
	SearchWithRequestOptions(query string, params Map, opts *RequestOptions) (res QueryRes, err error)

	// DeleteBy finds all the records that match the given query parameters
	// and deletes them. However, those parameters do not support all the
	// options of a query, only its filters (numeric, facet, or tag) and geo
	// queries. They also do not accept empty filters or query. More details
	// here:
	// https://www.algolia.com/doc/rest-api/search/#delete-by-query
	DeleteBy(params Map) (res UpdateTaskRes, err error)

	// DeleteByWithRequestOptions is the same as DeleteBy but it also accepts
	// extra RequestOptions.
	DeleteByWithRequestOptions(params Map, opts *RequestOptions) (res UpdateTaskRes, err error)

	// DeleteByQuery finds all the records that match the `query`, according to
	// the given 'params` and deletes them. It hangs until all the deletion
	// operations have completed.
	//
	// Deprecated: Use DeleteBy instead.
	DeleteByQuery(query string, params Map) error

	// DeleteByQueryWithRequestOptions is the same as DeleteByQuery but it also
	// accepts extra RequestOptions.
	//
	// Deprecated: Use DeleteByWithRequestOptions instead.
	DeleteByQueryWithRequestOptions(query string, params Map, opts *RequestOptions) error

	// SearchFacet searches inside a facet's values, optionally
	// restricting the returned values to those contained in objects matching
	// other (regular) search criteria. The `facet` parameter is the name of
	// the facet to search (must be declared in `attributesForFaceting`). The
	// `query` string is the text used to matched against facet's values. The
	// `params` controls the search parameters you want to apply against the
	// matching records. Note that it can be `nil` and that pagination
	// parameters are not taken into account.
	//
	// Deprecated: Use SearchForFacetValues instead.
	SearchFacet(facet, query string, params Map) (res SearchFacetRes, err error)

	// SearchForFacetValues searches inside a facet's values, optionally
	// restricting the returned values to those contained in objects matching
	// other (regular) search criteria. The `facet` parameter is the name of
	// the facet to search (must be declared in `attributesForFaceting`). The
	// `query` string is the text used to matched against facet's values. The
	// `params` controls the search parameters you want to apply against the
	// matching records. Note that it can be `nil` and that pagination
	// parameters are not taken into account.
	SearchForFacetValues(facet, query string, params Map) (res SearchFacetRes, err error)

	// SearchForFacetValuesWithRequestOptions is the same as
	// SearchForFacetValues but it also accepts extra RequestOptions.
	SearchForFacetValuesWithRequestOptions(facet, query string, params Map, opts *RequestOptions) (res SearchFacetRes, err error)

	// SaveRule saves the given Rule for the current index. If a Rule with the
	// same objectID already exists, it will get overriden. The operation can
	// be forwarded to the index replicas by setting `forwardToReplicas` to
	// `true`.
	SaveRule(rule Rule, forwardToReplicas bool) (SaveRuleRes, error)

	// SaveRuleWithRequestOptions is the same as SaveRule but it also accepts
	// extra RequestOptions.
	SaveRuleWithRequestOptions(rule Rule, forwardToReplicas bool, opts *RequestOptions) (SaveRuleRes, error)

	// SaveRule saves the given Rules by batch for the current index. The
	// operation can be forwarded to the index replicas by setting
	// `forwardToReplicas` to `true`. A `clear` operation can be applied before
	// writing the new Rules by setting `clearExistingRules` to `true`.
	BatchRules(rules []Rule, forwardToReplicas, clearExistingRules bool) (BatchRulesRes, error)

	// BatchRulesWithRequestOptions is the same as BatchRules but it also
	// accepts extra RequestOptions.
	BatchRulesWithRequestOptions(rules []Rule, forwardToReplicas, clearExistingRules bool, opts *RequestOptions) (BatchRulesRes, error)

	// GetRule returns the Rule identified by the given `objectID`. A non-nil
	// error is returned if the Rule cannot be found.
	GetRule(objectID string) (*Rule, error)

	// GetRuleWithRequestOptions is the same as GetRule but it also accepts
	// extra RequestOptions.
	GetRuleWithRequestOptions(objectID string, opts *RequestOptions) (*Rule, error)

	// DeleteRule deletes the Rule identified by the given `objectID` for the
	// current index. The operation can be forwarded to the index replicas by
	// setting `forwardToReplicas` to `true`.
	DeleteRule(objectID string, forwardToReplicas bool) (DeleteRuleRes, error)

	// DeleteRuleWithRequestOptions is the same as DeleteRule but it also
	// accepts extra RequestOptions.
	DeleteRuleWithRequestOptions(objectID string, forwardToReplicas bool, opts *RequestOptions) (DeleteRuleRes, error)

	// ClearRules removes all existing Rules for the current index. The
	// operation can be forwarded to the index replicas by setting
	// `forwardToReplicas` to `true`.
	ClearRules(forwardToReplicas bool) (ClearRulesRes, error)

	// ClearRulesWithRequestOptions is the same as ClearRules but it also
	// accepts extra RequestOptions.
	ClearRulesWithRequestOptions(forwardToReplicas bool, opts *RequestOptions) (ClearRulesRes, error)

	// SearchRules allows to search for Rules for the current index. The
	// given `Map` can be populated with any of the following fields, which are
	// all optional:
	//
	//   - `query` (string):                     enable full text search within the Rule fields
	//   - `anchoring` (RulePatternAnchoring):   restricts the search to Rules with a specific anchoring type
	//   - `context` (string):                   restricts the search to rules with a specific context (exact match)
	//   - `page` (int):                         requested page (zero-based, defaults to zero)
	//   - `hitsPerPage` (int):                  maximum number of hits in a page (defaults to 20)
	SearchRules(params Map) (SearchRulesRes, error)

	// SearchRulesWithRequestOptions is the same as SearchRules but it also
	// accepts extra RequestOptions.
	SearchRulesWithRequestOptions(params Map, opts *RequestOptions) (SearchRulesRes, error)

	// ReplaceAllSynonyms replace all the synonyms of the current index with the given ones.
	ReplaceAllSynonyms(synonyms []Synonym) (res UpdateTaskRes, err error)

	// ReplaceAllSynonymsWithRequestOptions is the same as ReplaceAllSynonyms but it also
	// accepts extra RequestOptions.
	ReplaceAllSynonymsWithRequestOptions(synonyms []Synonym, opts *RequestOptions) (res UpdateTaskRes, err error)

	// ReplaceAllRules replace all the rules of the current index with the given ones.
	ReplaceAllRules(rules []Rule) (res BatchRulesRes, err error)

	// ReplaceAllRulesWithRequestOptions is the same as ReplaceAllRules but it also
	// accepts extra RequestOptions.
	ReplaceAllRulesWithRequestOptions(rules []Rule, opts *RequestOptions) (res BatchRulesRes, err error)

	// ReplaceAllObjects replace all the objects of the current index with the given ones.
	ReplaceAllObjects(objects []Object) error

	// ReplaceAllObjectsWithRequestOptions is the same as ReplaceAllObjects but it also
	// accepts extra RequestOptions.
	ReplaceAllObjectsWithRequestOptions(objects []Object, opts *RequestOptions) error
}

// IndexIterator is used by the BrowseAll functions to iterate over all the
// records of an index (or a subset according to what the query and the params
// are).
type IndexIterator interface {
	// Next returns the next record each time is is called. Subsequent pages of
	// results are automatically loaded and an error is returned if a problem
	// occurs. When the last element is reached, an error is returned with the
	// following message: "No more hits".
	Next() (res Map, err error)
}

type Analytics interface {
	// AddABTest creates a new AB Test.
	AddABTest(abTest ABTest) (res ABTestTaskRes, err error)

	// DeleteABTest stops the AB Test referenced by the given ID.
	StopABTest(id int) (res ABTestTaskRes, err error)

	// DeleteABTest removes the AB Test referenced by the given ID.
	DeleteABTest(id int) (res ABTestTaskRes, err error)

	// GetABTest returns the informations relative to the AB Test referenced by
	// the given ID.
	GetABTest(id int) (res ABTestResponse, err error)

	// GetABTests retrieves a list of ABTests, according to the given
	// parameters. The returned list may not be exhaustive, depending on the
	// parameters that were provided.
	//
	// To retrieve the complete list of enabled AB tests, one should iterate
	// over the multiple pages of result returned by GetABTests until no more
	// AB Test is found.
	GetABTests(params Map) (res GetABTestsRes, err error)

	// WaitTask blocks until the given task has ended successfully. If anything
	// goes wrong or if the task did not succeed, a non-nil error is returned.
	WaitTask(task ABTestTaskRes) (err error)
}

type Insights interface {
	User(userToken string) InsightsWithUser
}

type InsightsWithUser interface {
	ClickedFilters(eventName, indexName string, filters []string) (res InsightsResponse, err error)
	ClickedFiltersWithRequestOptions(eventName, indexName string, filters []string, opts *RequestOptions) (res InsightsResponse, err error)
	ClickedObjectIDs(eventName, indexName string, objectIDs []string) (res InsightsResponse, err error)
	ClickedObjectIDsWithRequestOptions(eventName, indexName string, objectIDs []string, opts *RequestOptions) (res InsightsResponse, err error)
	ClickedObjectIDsAfterSearch(eventName, indexName string, objectIDs []string, positions []int, queryID string) (res InsightsResponse, err error)
	ClickedObjectIDsAfterSearchWithRequestOptions(eventName, indexName string, objectIDs []string, positions []int, queryID string, opts *RequestOptions) (res InsightsResponse, err error)

	ConvertedObjectIDs(eventName, indexName string, objectIDs []string) (res InsightsResponse, err error)
	ConvertedObjectIDsWithRequestOptions(eventName, indexName string, objectIDs []string, opts *RequestOptions) (res InsightsResponse, err error)
	ConvertedObjectIDsAfterSearch(eventName, indexName string, objectIDs []string, queryID string) (res InsightsResponse, err error)
	ConvertedObjectIDsAfterSearchWithRequestOptions(eventName, indexName string, objectIDs []string, queryID string, opts *RequestOptions) (res InsightsResponse, err error)
	ConvertedFilters(eventName, indexName string, filters []string) (res InsightsResponse, err error)
	ConvertedFiltersWithRequestOptions(eventName, indexName string, objectIDs []string, opts *RequestOptions) (res InsightsResponse, err error)

	ViewedFilters(eventName, indexName string, filters []string) (res InsightsResponse, err error)
	ViewedFiltersWithRequestOptions(eventName, indexName string, filters []string, opts *RequestOptions) (res InsightsResponse, err error)
	ViewedObjectIDs(eventName, indexName string, objectIDs []string) (res InsightsResponse, err error)
	ViewedObjectIDsWithRequestOptions(eventName, indexName string, objectIDs []string, opts *RequestOptions) (res InsightsResponse, err error)

	SendEvent(req InsightsRequest) (res InsightsResponse, err error)
	SendEventWithRequestOptions(req InsightsRequest, opts *RequestOptions) (res InsightsResponse, err error)
	SendEvents(req []InsightsRequest) (res InsightsResponse, err error)
	SendEventsWithRequestOptions(req []InsightsRequest, opts *RequestOptions) (res InsightsResponse, err error)
}

// AccountClient is responsible for handling cross-application operations.
type AccountClient interface {
	// CopyIndex copies the content of the entire source index to the destination index. Indices from the same
	// application cannot be copied. To do so, use Client.CopyIndex instead.
	CopyIndex(src, dst Index) (taskIDs []int, err error)

	// CopyIndexWithRequestOptions is the same as CopyIndex but it also
	// accepts extra RequestOptions.
	CopyIndexWithRequestOptions(src, dst Index, opts *RequestOptions) (taskIDs []int, err error)
}
