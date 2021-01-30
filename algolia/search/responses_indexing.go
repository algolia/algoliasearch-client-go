package search

import (
	"fmt"
	"sync"
	"time"
)

type SaveObjectRes struct {
	CreatedAt time.Time `json:"createdAt"`
	ObjectID  string    `json:"objectID"`
	TaskID    int64     `json:"taskID"`
	wait      func(taskID int64, opts ...interface{}) error
}

func (r SaveObjectRes) Wait(opts ...interface{}) error {
	return r.wait(r.TaskID, opts...)
}

type BatchRes struct {
	ObjectIDs []string `json:"objectIDs"`
	TaskID    int64    `json:"taskID"`
	wait      func(taskID int64, opts ...interface{}) error
}

func (r BatchRes) Wait(opts ...interface{}) error {
	return r.wait(r.TaskID, opts...)
}

type GroupBatchRes struct {
	Responses []BatchRes
}

func (r GroupBatchRes) Wait(opts ...interface{}) error {
	var wg sync.WaitGroup
	errs := make(chan error, len(r.Responses))

	for _, res := range r.Responses {
		wg.Add(1)
		go func(wg *sync.WaitGroup, res BatchRes) {
			errs <- res.Wait(opts...)
			wg.Done()
		}(&wg, res)
	}

	go func() {
		wg.Wait()
		close(errs)
	}()

	for err := range errs {
		if err != nil {
			return fmt.Errorf("at least one batch could not complete: %s", err)
		}
	}

	return nil
}

func (r GroupBatchRes) ObjectIDs() []string {
	var objectIDs []string
	for _, res := range r.Responses {
		objectIDs = append(objectIDs, res.ObjectIDs...)
	}
	return objectIDs
}
