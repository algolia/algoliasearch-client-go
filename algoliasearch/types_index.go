package algoliasearch

type IndexRes struct {
	CreatedAt      string `json:"createdAt"`
	DataSize       int64  `json:"dataSize"`
	Entries        int64  `json:"entries"`
	LastBuildTimeS int64  `json:"lastBuildTimeS"`
	Name           string `json:"name"`
	PendingTask    bool   `json:"pendingTask"`
	UpdatedAt      string `json:"updatedAt"`
}

type listIndexesRes struct {
	items []IndexRes
}
