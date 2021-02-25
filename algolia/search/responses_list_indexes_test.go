package search

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestListIndexesResponse(t *testing.T) {
	for _, tc := range []struct {
		jsonResponse string
		primary      string
		replicas     []string
	}{
		{
			`{
				"items":[
				{
					"name": "airports",
					"createdAt": "2017-05-15T08:09:45.173Z",
					"updatedAt": "2017-05-15T08:09:45.173Z",
					"entries":  1005,
					"dataSize": 0,
					"fileSize": 0,
					"lastBuildTimeS": 1,
					"numberOfPendingTasks": 0,
					"pendingTask": false,
					"replicas": ["replica_index"]
				}
				],
				"nbPages":1
			}`,
			"",
			[]string{"replica_index"},
		},
		{
			`{
				"items":[
				{
					"name": "airports",
					"createdAt": "2017-05-15T08:09:45.173Z",
					"updatedAt": "2017-05-15T08:09:45.173Z",
					"entries":  1005,
					"dataSize": 0,
					"fileSize": 0,
					"lastBuildTimeS": 1,
					"numberOfPendingTasks": 0,
					"pendingTask": false,
					"primary": "primary_index"
				}
				],
				"nbPages":1
			}`,
			"primary_index",
			nil,
		},
	} {
		var listIndicesRes ListIndicesRes
		err := json.Unmarshal([]byte(tc.jsonResponse), &listIndicesRes)
		require.NoError(t, err)
		require.Equal(t, 1, len(listIndicesRes.Items))
		require.Equal(t, 1, listIndicesRes.NbPages)
		require.Equal(t, "airports", listIndicesRes.Items[0].Name)
		require.Equal(t, "2017-05-15 08:09:45.173 +0000 UTC", listIndicesRes.Items[0].CreatedAt.String())
		require.Equal(t, "2017-05-15 08:09:45.173 +0000 UTC", listIndicesRes.Items[0].UpdatedAt.String())
		require.Equal(t, int64(1005), listIndicesRes.Items[0].Entries)
		require.Equal(t, int64(0), listIndicesRes.Items[0].DataSize)
		require.Equal(t, int64(0), listIndicesRes.Items[0].FileSize)
		require.Equal(t, int64(0), listIndicesRes.Items[0].NumberOfPendingTasks)
		require.Equal(t, false, listIndicesRes.Items[0].PendingTask)
		require.Equal(t, tc.primary, listIndicesRes.Items[0].Primary)
		require.Equal(t, tc.replicas, listIndicesRes.Items[0].Replicas)
	}
}
