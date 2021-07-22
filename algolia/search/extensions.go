package search

// Index extensions container
type Extensions struct {
	// Reranking extension configuration
	Reranking *Reranking `json:"reranking,omitempty"`
}

// Reranking extension settings
type Reranking struct {
	// Whether the extension is enabled
	Enabled bool `json:"enabled,omitempty"`

	// Amount of hits to rerank
	MaxNbHits int `json:"maxNbHits,omitempty"`

	// URL to target
	Endpoint string `json:"endpoint,omitempty"`
}
