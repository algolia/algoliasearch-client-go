package algoliasearch

type CreateObjectRes struct {
	CreatedAt string `json:"createdAt"`
	ObjectID  string `json:"objectID"`
	TaskID    int  `json:"taskID"`
}

type UpdateObjectRes struct {
	ObjectID  string `json:"objectID"`
	TaskID    int  `json:"taskID"`
	UpdatedAt string `json:"updatedAt"`
}

type DeleteRes struct {
	DeletedAt string `json:"deletedAt"`
}
