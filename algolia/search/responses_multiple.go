package search

import (
	"fmt"
)

type MultipleBatchRes struct {
	ObjectIDs []string         `json:"objectIDs"`
	TaskIDs   map[string]int64 `json:"taskID"`
	wait      func(index string, taskID int64, opts ...interface{}) error
}

func (r MultipleBatchRes) Wait(opts ...interface{}) error {
	errs := make(chan error, len(r.TaskIDs))

	for index, taskID := range r.TaskIDs {
		errs <- r.wait(index, taskID, opts...)
	}

	close(errs)

	for err := range errs {
		if err != nil {
			return fmt.Errorf("at least one batch could not complete: %s", err)
		}
	}

	return nil
}

type getObjectsRes struct {
	Results interface{} `json:"results"`
}

type MultipleQueriesRes struct {
	Results []IndexedQueryRes `json:"results"`
}

type IndexedQueryRes struct {
	Processed bool `json:"processed"`
	QueryRes
}
