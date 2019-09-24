package search

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/errs"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/opt"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/search"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/wait"
	"github.com/algolia/algoliasearch-client-go/v3/cts"
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

func TestSecuredAPIKeyShouldNotBeExpired(t *testing.T) {
	c := search.NewClient("", "")
	searchKey := os.Getenv("ALGOLIA_SEARCH_KEY_1")
	require.NotEmpty(t, searchKey)

	securedKey, err := search.GenerateSecuredAPIKey(
		searchKey,
		opt.ValidUntil(time.Now().Add(10*time.Minute)),
		opt.RestrictIndices("indexName"),
	)
	require.NoError(t, err)
	validity, err := c.GetSecuredAPIKeyRemainingValidity(securedKey)
	require.NoError(t, err)
	require.Greater(t, int(validity.Seconds()), 0)
}

func TestSecuredAPIKeyShouldBeExpired(t *testing.T) {
	c := search.NewClient("", "")
	searchKey := os.Getenv("ALGOLIA_SEARCH_KEY_1")
	require.NotEmpty(t, searchKey)

	securedKey, err := search.GenerateSecuredAPIKey(
		searchKey,
		opt.ValidUntil(time.Now().Add(-10*time.Minute)),
		opt.RestrictIndices("indexName"),
	)
	require.NoError(t, err)
	validity, err := c.GetSecuredAPIKeyRemainingValidity(securedKey)
	require.NoError(t, err)
	require.LessOrEqual(t, int(validity.Seconds()), 0)
}

func TestSecuredAPIKeyParametersValidity(t *testing.T) {
	c := search.NewClient("", "")
	searchKey := os.Getenv("ALGOLIA_SEARCH_KEY_1")
	require.NotEmpty(t, searchKey)

	securedKey, err := search.GenerateSecuredAPIKey(
		searchKey,
		opt.RestrictIndices("indexName"),
	)
	require.NoError(t, err)
	_, err = c.GetSecuredAPIKeyRemainingValidity(securedKey)
	require.EqualError(t, err, errs.ErrValidUntilNotFound.Error())
}
