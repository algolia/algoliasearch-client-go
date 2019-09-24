package search

import (
	"os"
	"testing"
	"time"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/opt"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/search"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/wait"
	"github.com/algolia/algoliasearch-client-go/v3/cts"
	"github.com/stretchr/testify/require"
)

func TestSecuredAPIKeys(t *testing.T) {
	t.Parallel()

	client, index1, indexName1 := cts.InitSearchClient1AndIndex(t)
	indexName2 := indexName1 + "_dev"
	index2 := client.InitIndex(indexName2)

	// Create the two indices by adding a dummy object to each
	{
		g := wait.NewGroup()
		obj := map[string]string{"objectID": "one"}

		res, err := index1.SaveObject(obj)
		require.NoError(t, err)
		g.Collect(res)

		res, err = index2.SaveObject(obj)
		require.NoError(t, err)
		g.Collect(res)

		require.NoError(t, g.Wait())
	}

	// Generate the key
	var generatedKey string
	{
		sourceKey := os.Getenv("ALGOLIA_SEARCH_KEY_1")
		require.NotEmpty(t, sourceKey)

		var err error
		generatedKey, err = search.GenerateSecuredAPIKey(sourceKey,
			opt.ValidUntil(time.Now().Add(10*time.Minute)),
			opt.RestrictIndices(indexName1),
		)
		require.NoError(t, err)
	}

	// Try to use the key on authorized and restricted indices
	{
		appID := os.Getenv("ALGOLIA_APPLICATION_ID_1")
		client := search.NewClient(appID, generatedKey)

		_, err := client.InitIndex(indexName1).Search("")
		require.NoError(t, err)

		_, err = client.InitIndex(indexName2).Search("")
		require.Error(t, err)
	}
}
