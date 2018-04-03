package algoliasearch

type IndexRes struct {
	CreatedAt            string `json:"createdAt"`
	DataSize             int    `json:"dataSize"`
	Entries              int    `json:"entries"`
	FileSize             int    `json:"fileSize"`
	LastBuildTimeS       int    `json:"lastBuildTimeS"`
	Name                 string `json:"name"`
	NumberOfPendingTasks int    `json:"numberOfPendingTasks"`
	PendingTask          bool   `json:"pendingTask"`
	UpdatedAt            string `json:"updatedAt"`
}

type listIndexesRes struct {
	Items   []IndexRes `json:"items"`
	NbPages int        `json:"nbPages"`
}
