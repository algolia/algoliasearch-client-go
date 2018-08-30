package search_client

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/algolia/algoliasearch-client-go/algoliasearch"
	"github.com/algolia/algoliasearch-client-go/it"
)

func TestSecuredAPIKeys(t *testing.T) {
	t.Parallel()

	client, index1, indexName1 := it.InitSearchClient1AndIndex(t)
	indexName2 := indexName1 + "_dev"
	index2 := client.InitIndex(indexName2)

	// Create the two indices by adding a dummy object to each
	{
		obj := algoliasearch.Object{"objectID": "one"}

		res, err := index1.AddObject(obj)
		require.NoError(t, err)
		it.WaitTasks(t, index1, res.TaskID)

		res, err = index2.AddObject(obj)
		require.NoError(t, err)
		it.WaitTasks(t, index2, res.TaskID)
	}

	// Generate the key
	var generatedKey string
	{
		sourceKey := os.Getenv("ALGOLIA_SEARCH_KEY_1")
		require.NotEmpty(t, sourceKey)

		params := algoliasearch.Map{
			"validUntil":      int(time.Now().Add(10 * time.Minute).Unix()),
			"restrictIndices": indexName1,
		}

		var err error
		generatedKey, err = algoliasearch.GenerateSecuredAPIKey(sourceKey, params)
		require.NoError(t, err)
	}

	// Try to use the key on authorized and restricted indices
	{
		appID := os.Getenv("ALGOLIA_APPLICATION_ID_1")
		client := algoliasearch.NewClient(appID, generatedKey)

		_, err := client.InitIndex(indexName1).Search("", nil)
		require.NoError(t, err)

		_, err = client.InitIndex(indexName2).Search("", nil)
		require.Error(t, err)
	}
}
