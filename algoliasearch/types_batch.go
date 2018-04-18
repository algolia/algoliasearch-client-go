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
	TaskID    int      `json:"taskID"`
}

type MultipleBatchRes struct {
	ObjectIDs []string       `json:"objectIDs"`
	TaskID    map[string]int `json:"taskID"`
}

func newBatchOperations(objects []Object, action string) (operations []BatchOperation, err error) {
	operations = make([]BatchOperation, len(objects))

	for i, o := range objects {
		// In the case of something else than `addObject` and `clear`
		// operations, the `objectID` field is required.
		if action != "addObject" && action != "clear" {
			_, err = o.ObjectID()
			if err != nil {
				return
			}
		}

		operations[i].Action = action
		operations[i].Body = o
	}

	return
}
