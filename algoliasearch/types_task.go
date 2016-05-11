package algoliasearch

type DeleteTaskRes struct {
	DeletedAt string `json:"deletedAt"`
	TaskID    int64  `json:"taskID"`
}

type UpdateTaskRes struct {
	TaskID    int64  `json:"taskID"`
	UpdatedAt string `json:"updatedAt"`
}

type TaskStatusRes struct {
	Status      string `json:"status"`
	PendingTask bool   `json:"pendingTask"`
}
