package algoliasearch

type BatchOperation struct {
	Action string      `json:"action"`
	Body   interface{} `json:"body,omitempty"`
}

type BatchOperationIndexed struct {
	BatchOperation
	IndexName string `json:"indexName"`
}

type BatchRes struct {
	ObjectIDs []string `json:"objectIDs"`
	TaskID    int    `json:"taskID"`
}

type MultipleBatchRes struct {
	ObjectIDs []string         `json:"objectIDs"`
	TaskID    map[string]int `json:"taskID"`
}

func newBatchOperations(objects []Object, action string) []BatchOperation {
	operations := make([]BatchOperation, len(objects))

	for i, o := range objects {
		operations[i].Action = action
		operations[i].Body = o
	}

	return operations
}
