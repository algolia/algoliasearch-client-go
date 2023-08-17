package search

import (
	"time"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/call"
)

type ClientInterface interface {
	// Misc
	WaitTask(taskID int64, opts ...interface{}) error
	WaitRecommendTask(taskID int64, opts ...interface{}) error
	GetStatus(taskID int64, opts ...interface{}) (res TaskStatusRes, err error)
	GetRecommendStatus(taskID int64, opts ...interface{}) (res TaskStatusRes, err error)
	InitIndex(indexName string) *Index
	ListIndices(opts ...interface{}) (res ListIndicesRes, err error)
	GetLogs(opts ...interface{}) (res GetLogsRes, err error)
	CustomRequest(res interface{}, method string, path string, body interface{}, k call.Kind, opts ...interface{}) error

	// Copy index operations
	CopyRules(source, destination string, opts ...interface{}) (UpdateTaskRes, error)
	CopySettings(source, destination string, opts ...interface{}) (UpdateTaskRes, error)
	CopySynonyms(source, destination string, opts ...interface{}) (UpdateTaskRes, error)
	CopyIndex(source, destination string, opts ...interface{}) (UpdateTaskRes, error)

	// Move index operations
	MoveIndex(source, destination string, opts ...interface{}) (UpdateTaskRes, error)

	// API key methods
	GetAPIKey(keyID string, opts ...interface{}) (key Key, err error)
	AddAPIKey(key Key, opts ...interface{}) (res CreateKeyRes, err error)
	UpdateAPIKey(key Key, opts ...interface{}) (res UpdateKeyRes, err error)
	DeleteAPIKey(keyID string, opts ...interface{}) (res DeleteKeyRes, err error)
	RestoreAPIKey(keyID string, opts ...interface{}) (res RestoreKeyRes, err error)
	ListAPIKeys(opts ...interface{}) (res ListAPIKeysRes, err error)
	GetSecuredAPIKeyRemainingValidity(keyID string, opts ...interface{}) (v time.Duration, err error)

	// Multiple methods
	MultipleBatch(operations []BatchOperationIndexed, opts ...interface{}) (res MultipleBatchRes, err error)
	MultipleGetObjects(requests []IndexedGetObject, objects interface{}, opts ...interface{}) (err error)
	MultipleQueries(queries []IndexedQuery, strategy string, opts ...interface{}) (res MultipleQueriesRes, err error)

	// Multi-Cluster Management (MCM) methods
	ListClusters(opts ...interface{}) (res ListClustersRes, err error)
	ListUserIDs(opts ...interface{}) (res ListUserIDsRes, err error)
	GetUserID(userID string, opts ...interface{}) (res UserID, err error)
	AssignUserID(userID, clusterName string, opts ...interface{}) (res AssignUserIDRes, err error)
	AssignUserIDs(userIDs []string, clusterName string, opts ...interface{}) (res AssignUserIDRes, err error)
	RemoveUserID(userID string, opts ...interface{}) (res RemoveUserIDRes, err error)
	GetTopUserIDs(opts ...interface{}) (res TopUserIDs, err error)
	SearchUserIDs(query string, opts ...interface{}) (res SearchUserIDRes, err error)
	HasPendingMappings(opts ...interface{}) (res HasPendingMappingsRes, err error)

	// Custom dictionaries methods
	SaveDictionaryEntries(dictionaryName DictionaryName, dictionaryEntries []DictionaryEntry, opts ...interface{}) (res UpdateTaskRes, err error)
	ReplaceDictionaryEntries(dictionaryName DictionaryName, dictionaryEntries []DictionaryEntry, opts ...interface{}) (res UpdateTaskRes, err error)
	DeleteDictionaryEntries(dictionaryName DictionaryName, objectIDs []string, opts ...interface{}) (res UpdateTaskRes, err error)
	ClearDictionaryEntries(dictionaryName DictionaryName, opts ...interface{}) (res UpdateTaskRes, err error)
	SearchDictionaryEntries(dictionaryName DictionaryName, query string, opts ...interface{}) (res SearchDictionariesRes, err error)
	GetDictionarySettings(opts ...interface{}) (res DictionarySettings, err error)
	SetDictionarySettings(settings DictionarySettings, opts ...interface{}) (res UpdateTaskRes, err error)

	// Personalization

	// Deprecated: use recommendation.Client.SetPersonalizationStrategy() instead
	SetPersonalizationStrategy(strategy Strategy, opts ...interface{}) (res SetPersonalizationStrategyRes, err error)
	// Deprecated: use recommendation.Client.GetPersonalizationStrategy() instead
	GetPersonalizationStrategy(opts ...interface{}) (res GetPersonalizationStrategyRes, err error)
}

var _ ClientInterface = &Client{}
