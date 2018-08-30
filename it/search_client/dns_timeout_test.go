package search_client

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/algolia/algoliasearch-client-go/algoliasearch"
	"github.com/algolia/algoliasearch-client-go/it"
)

func TestDNSTimeout(t *testing.T) {
	t.Parallel()

	appID, apiKey := it.GetTestingCredentials(t, "ALGOLIA_APPLICATION_ID_1", "ALGOLIA_ADMIN_KEY_1")

	client := algoliasearch.NewClientWithHosts(
		appID,
		apiKey,
		[]string{
			"algolia.biz",
			appID + "-1.algolianet.com",
			appID + "-2.algolianet.com",
			appID + "-2.algolianet.com",
		},
	)

	start := time.Now()
	for i := 0; i < 10; i++ {
		_, err := client.ListIndexes()
		require.NoError(t, err)
	}
	delta := time.Since(start)
	expectedMax := 5 * time.Second
	require.True(t, delta < expectedMax, "retries should have taken less than %s but took %s", expectedMax, delta)
}
