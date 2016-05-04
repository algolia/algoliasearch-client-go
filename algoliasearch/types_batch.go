package algoliasearch

type CustomBatchRes struct {
	ObjectIDs []string `json:"objectIDs"`
	TaskID    int64    `json:"taskID"`
}

type BatchRecord struct {
	Action    string                 `json:"action"`
	Body      map[string]interface{} `json:"body"`
	IndexName string                 `json:"indexName"`
}
