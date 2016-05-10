package algoliasearch

type CreateObjectRes struct {
	CreatedAt string `json:"createdAt"`
	ObjectID  string `json:"objectID"`
	TaskID    int64  `json:"taskID"`
}

type UpdateObjectRes struct {
	ObjectID  string `json:"objectID"`
	TaskID    int64  `json:"taskID"`
	UpdatedAt string `json:"updatedAt"`
}

type DeleteRes struct {
	DeletedAt string `json:"deletedAt"`
}

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

type BatchRes struct {
	ObjectIDs []string `json:"objectIDs"`
	TaskID    int64    `json:"taskID"`
}
