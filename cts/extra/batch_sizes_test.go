package extra

import (
	"fmt"
	"math"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/search"
	"github.com/algolia/algoliasearch-client-go/v3/cts"
)

func TestBatchSizes(t *testing.T) {
	initIndexWithMaxBatchSize := func(maxBatchSize int) *search.Index {
		appID, apiKey := cts.GetTestingCredentials(t, "ALGOLIA_APPLICATION_ID_1", "ALGOLIA_ADMIN_KEY_1")
		indexName := cts.GenerateIndexName(t)
		client := search.NewClientWithConfig(search.Configuration{
			AppID:        appID,
			APIKey:       apiKey,
			MaxBatchSize: maxBatchSize,
		})
		return client.InitIndex(indexName)
	}

	for _, c := range []struct {
		batchSize    int
		maxBatchSize int
	}{
		{0, 1000},
		{1, 1000},
		{10, 1000},
		{999, 1000},
		{1000, 1000},
		{1001, 1000},
		{1999, 1000},
		{2000, 1000},
		{2001, 1000},
		{100, 9},
	} {
		var (
			expectedNbBatches        int
			expectedRecordsPerBatch  int
			expectedRecordsLastBatch int
			batch                    []map[string]string
		)

		if c.batchSize == 0 {
			expectedNbBatches = 0
			expectedRecordsPerBatch = 0
			expectedRecordsLastBatch = 0
		} else if c.batchSize <= c.maxBatchSize {
			expectedNbBatches = 1
			expectedRecordsPerBatch = 0
			expectedRecordsLastBatch = c.batchSize
		} else {
			expectedNbBatches = int(math.Ceil(float64(c.batchSize) / float64(c.maxBatchSize)))
			expectedRecordsPerBatch = c.maxBatchSize
			remaining := c.batchSize % c.maxBatchSize
			if remaining == 0 {
				expectedRecordsLastBatch = c.maxBatchSize
			} else {
				expectedRecordsLastBatch = remaining
			}
		}

		for i := 0; i < c.batchSize; i++ {
			batch = append(batch, map[string]string{"objectID": strconv.Itoa(i)})
		}

		name := fmt.Sprintf("trying to save %d records in %d batches", c.batchSize, expectedNbBatches)
		index := initIndexWithMaxBatchSize(c.maxBatchSize)
		res, err := index.SaveObjects(batch)

		require.NoError(t, err, name)
		require.Len(t, res.ObjectIDs(), c.batchSize, name)
		require.Len(t, res.Responses, expectedNbBatches, name)

		for i := 0; i < expectedNbBatches-1; i++ {
			require.Len(t, res.Responses[i].ObjectIDs, expectedRecordsPerBatch, name)
		}
		if expectedNbBatches > 0 {
			require.Len(t, res.Responses[expectedNbBatches-1].ObjectIDs, expectedRecordsLastBatch, name)
		}
	}
}
