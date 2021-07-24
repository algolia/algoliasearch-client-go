package search

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnmarshalRerankingExtension(t *testing.T) {

	payload := `{
		"reranking": {
			"enabled": true,
			"maxNbHits": 100,
			"endpoint": "example.org/rerankingextension"
		}
	}`

	var e Extensions
	err := json.Unmarshal([]byte(payload), &e)
	require.NoError(t, err)
	require.NotNil(t, e.Reranking)
	require.Equal(t, e.Reranking.Enabled, true)
	require.Equal(t, e.Reranking.Endpoint, "example.org/rerankingextension")
	require.Equal(t, e.Reranking.MaxNbHits, 100)

}
