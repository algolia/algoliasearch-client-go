package algoliasearch

type IndexRes struct {
	CreatedAt      string `json:"createdAt"`
	DataSize       int    `json:"dataSize"`
	Entries        int    `json:"entries"`
	LastBuildTimeS int    `json:"lastBuildTimeS"`
	Name           string `json:"name"`
	PendingTask    bool   `json:"pendingTask"`
	UpdatedAt      string `json:"updatedAt"`
}

type listIndexesRes struct {
	Items []IndexRes
}
