package search

import "github.com/algolia/algoliasearch-client-go/algolia/call"

type ClientInterface interface {
	// Misc
	InitIndex(indexName string) *Index
	ListIndexes(opts ...interface{}) (res ListIndexesRes, err error)
	GetLogs(opts ...interface{}) (res GetLogsRes, err error)
	CustomRequest(res interface{}, method string, path string, body interface{}, k call.Kind, opts ...interface{}) error

	// Copy index operations
	CopyRules(source, destination string, opts ...interface{}) (UpdateTaskRes, error)
	CopySettings(source, destination string, opts ...interface{}) (UpdateTaskRes, error)
	CopySynonyms(source, destination string, opts ...interface{}) (UpdateTaskRes, error)
	CopyIndex(source, destination string, opts ...interface{}) (UpdateTaskRes, error)

	// Move index operations
	MoveRules(source, destination string, opts ...interface{}) (UpdateTaskRes, error)
	MoveSettings(source, destination string, opts ...interface{}) (UpdateTaskRes, error)
	MoveSynonyms(source, destination string, opts ...interface{}) (UpdateTaskRes, error)
	MoveIndex(source, destination string, opts ...interface{}) (UpdateTaskRes, error)

	// API key methods
	GetAPIKey(keyID string, opts ...interface{}) (key Key, err error)
	AddAPIKey(key Key, opts ...interface{}) (res CreateKeyRes, err error)
	UpdateAPIKey(key Key, opts ...interface{}) (res UpdateKeyRes, err error)
	DeleteAPIKey(keyID string, opts ...interface{}) (res DeleteKeyRes, err error)
	RestoreAPIKey(keyID string, opts ...interface{}) (res RestoreKeyRes, err error)
	ListAPIKeys(opts ...interface{}) (res ListAPIKeysRes, err error)

	// Multiple methods
	MultipleBatch(operations []BatchOperationIndexed, opts ...interface{}) (res MultipleBatchRes, err error)
	MultipleGetObjects(requests []IndexedGetObject, objects interface{}, opts ...interface{}) (err error)
	MultipleQueries(queries []IndexedQuery, strategy string, opts ...interface{}) (res MultipleQueriesRes, err error)

	// Multi-Cluster Management (MCM) methods
	ListClusters(opts ...interface{}) (res ListClustersRes, err error)
	ListUserIDs(opts ...interface{}) (res ListUserIDsRes, err error)
	GetUserID(userID string, opts ...interface{}) (res UserID, err error)
	AssignUserID(userID, clusterName string, opts ...interface{}) (res AssignUserIDRes, err error)
	RemoveUserID(userID string, opts ...interface{}) (res RemoveUserIDRes, err error)
	GetTopUserIDs(opts ...interface{}) (res TopUserIDs, err error)
	SearchUserIDs(query string, opts ...interface{}) (res SearchUserIDRes, err error)

	// Personalization
	SetPersonalizationStrategy(strategy Strategy, opts ...interface{}) (res SetPersonalizationStrategyRes, err error)
	GetPersonalizationStrategy(opts ...interface{}) (res GetPersonalizationStrategyRes, err error)
}

var _ ClientInterface = &Client{}
