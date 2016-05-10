package algoliasearch

type BatchOperation struct {
	Action string      `json:"action"`
	Body   interface{} `json:"body,omitempty"`
}

func newBatchOperations(objects []Object, action string) []BatchOperation {
	operations := make([]BatchOperation, len(objects))

	for i, o := range objects {
		operations[i].Action = action
		operations[i].Body = o
	}

	return operations
}
