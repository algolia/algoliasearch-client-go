package analytics

import (
	"strings"
	"testing"
	"time"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/analytics"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/opt"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/search"
	"github.com/algolia/algoliasearch-client-go/v3/cts"
	"github.com/stretchr/testify/require"
)

func TestAATesting(t *testing.T) {
	t.Parallel()
	_, index, indexName := cts.InitSearchClient1AndIndex(t)
	analyticsClient := cts.InitAnalyticsClient1(t)

	// Remove old AB tests
	{
		var toRemove []int
		today := cts.TodayDate()

		res, err := analyticsClient.GetABTests()
		require.NoError(t, err)
		for _, abTest := range res.ABTests {
			if strings.HasPrefix("go-", abTest.Name) &&
				!strings.HasPrefix("go-"+today, abTest.Name) {
				toRemove = append(toRemove, abTest.ABTestID)
			}
		}

		for _, id := range toRemove {
			_, _ = analyticsClient.DeleteABTest(id)
		}
	}

	// Add a dummy object to the index
	{
		res, err := index.SaveObject(map[string]string{"objectID": "one"})
		require.NoError(t, err)
		require.NoError(t, res.Wait())
	}

	var abTestID int
	now := time.Now()

	abTest := analytics.ABTest{
		Name: cts.GenerateCanonicalPrefixName(),
		Variants: []analytics.Variant{
			{
				Index:             indexName,
				TrafficPercentage: 90,
			},
			{
				Index:             indexName,
				TrafficPercentage: 10,
				CustomSearchParameters: &search.QueryParams{
					IgnorePlurals: opt.IgnorePlurals(true),
				},
			},
		},
		EndAt: now.Truncate(time.Hour).Add(24 * time.Hour),
	}

	// Create the AB test
	{
		res, err := analyticsClient.AddABTest(abTest)
		require.NoError(t, err)
		require.NoError(t, res.Wait())
		abTestID = res.ABTestID
	}

	// Retrieve the AB test and check it corresponds to the original one
	{
		res, err := analyticsClient.GetABTest(abTestID)
		require.NoError(t, err)
		checkABTestsAreEqual(t, abTest, res)
		require.NotEqual(t, res.Status, "stopped")
	}
}
