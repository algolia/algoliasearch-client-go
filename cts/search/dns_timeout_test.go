package search

import (
	"testing"
	"time"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/search"
	"github.com/algolia/algoliasearch-client-go/v3/cts"
	"github.com/stretchr/testify/require"
)

func TestDNSTimeout(t *testing.T) {
	t.Parallel()

	appID, apiKey := cts.GetTestingCredentials(t, "ALGOLIA_APPLICATION_ID_1", "ALGOLIA_ADMIN_KEY_1")

	client := search.NewClientWithConfig(
		search.Configuration{
			AppID:  appID,
			APIKey: apiKey,
			Hosts: []string{
				"algolia.biz",
				appID + "-1.algolianet.com",
				appID + "-2.algolianet.com",
				appID + "-2.algolianet.com",
			},
		},
	)

	start := time.Now()
	for i := 0; i < 10; i++ {
		_, err := client.ListIndices()
		require.NoError(t, err)
	}
	delta := time.Since(start)
	expectedMax := 5 * time.Second
	require.True(t, delta < expectedMax, "retries should have taken less than %s but took %s", expectedMax, delta)
}
