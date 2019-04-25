package insights

type UserClientInterface interface {
	ClickedObjectIDs(eventName string, indexName string, objectIDs []string, opts ...interface{}) (StatusMessageRes, error)
	ClickedObjectIDsAfterSearch(eventName string, indexName string, objectIDs []string, positions []int, queryID string, opts ...interface{}) (StatusMessageRes, error)
	ClickedFilters(eventName string, indexName string, filters []string, opts ...interface{}) (res StatusMessageRes, err error)
	ConvertedObjectIDs(eventName string, indexName string, objectIDs []string, opts ...interface{}) (StatusMessageRes, error)
	ConvertedObjectIDsAfterSearch(eventName string, indexName string, objectIDs []string, queryID string, opts ...interface{}) (StatusMessageRes, error)
	ConvertedFilters(eventName string, indexName string, filters []string, opts ...interface{}) (StatusMessageRes, error)
	ViewedObjectIDs(eventName string, indexName string, objectIDs []string, opts ...interface{}) (StatusMessageRes, error)
	ViewedFilters(eventName string, indexName string, filters []string, opts ...interface{}) (StatusMessageRes, error)
}
