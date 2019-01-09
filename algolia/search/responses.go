package search

import (
	"fmt"
	"sync"
)

type TaskStatusRes struct {
	Status      string `json:"status"`
	PendingTask bool   `json:"pendingTask"`
}

type UpdateTaskRes struct {
	TaskID    int    `json:"taskID"`
	UpdatedAt string `json:"updatedAt"`
	wait      func(taskID int) error
}

func (r UpdateTaskRes) Wait() error {
	return r.wait(r.TaskID)
}

type SaveObjectRes struct {
	CreatedAt string `json:"createdAt"`
	ObjectID  string `json:"objectID"`
	TaskID    int    `json:"taskID"`
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

type MultipleBatchRes struct {
	responses []BatchRes
}

func (r MultipleBatchRes) Wait() error {
	var wg sync.WaitGroup
	errs := make(chan error, len(r.responses))

	for _, res := range r.responses {
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
