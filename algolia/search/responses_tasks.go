package search

import "time"

type TaskStatusRes struct {
	Status      string `json:"status"`
	PendingTask bool   `json:"pendingTask"`
}

type UpdateTaskRes struct {
	TaskID    int64     `json:"taskID"`
	UpdatedAt time.Time `json:"updatedAt"`
	wait      func(taskID int64) error
}

type DeleteTaskRes struct {
	DeletedAt time.Time `json:"deletedAt"`
	TaskID    int64     `json:"taskID"`
	wait      func(taskID int64) error
}

func (r UpdateTaskRes) Wait() error { return r.wait(r.TaskID) }
func (r DeleteTaskRes) Wait() error { return r.wait(r.TaskID) }
