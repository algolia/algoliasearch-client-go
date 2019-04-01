package search

import (
	"fmt"
	"sync"
	"time"
)

type SaveObjectRes struct {
	CreatedAt time.Time `json:"createdAt"`
	ObjectID  string    `json:"objectID"`
	TaskID    int       `json:"taskID"`
	wait      func(taskID int) error
}

func (r SaveObjectRes) Wait() error {
	return r.wait(r.TaskID)
}

type BatchRes struct {
	ObjectIDs []string `json:"objectIDs"`
	TaskID    int      `json:"taskID"`
	wait      func(taskID int) error
}

func (r BatchRes) Wait() error {
	return r.wait(r.TaskID)
}

type GroupBatchRes struct {
	Responses []BatchRes
}

func (r GroupBatchRes) Wait() error {
	var wg sync.WaitGroup
	errs := make(chan error, len(r.Responses))

	for _, res := range r.Responses {
		wg.Add(1)
		go func(wg *sync.WaitGroup, res BatchRes) {
			errs <- res.Wait()
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
