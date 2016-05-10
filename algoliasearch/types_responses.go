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

type BatchRes struct {
	ObjectIDs []string `json:"objectIDs"`
	TaskID    int64    `json:"taskID"`
}
