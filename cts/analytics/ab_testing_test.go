package analytics

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/analytics"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/opt"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/wait"
	"github.com/algolia/algoliasearch-client-go/v3/cts"
)

func TestABTesting(t *testing.T) {
	t.Parallel()
	searchClient, index1, indexName1 := cts.InitSearchClient1AndIndex(t)
	indexName2 := indexName1 + "_dev"
	index2 := searchClient.InitIndex(indexName2)
	analyticsClient := cts.InitAnalyticsClient1(t)

	// Create the two indices by adding a dummy object in each of them
	{
		g := wait.NewGroup()

		res, err := index1.SaveObject(map[string]string{"objectID": "one"})
		require.NoError(t, err)
		g.Collect(res)

		res, err = index2.SaveObject(map[string]string{"objectID": "one"})
		require.NoError(t, err)
		g.Collect(res)

		require.NoError(t, g.Wait())
	}
	// Wait to make sure the indices are propagated to all nodes not only the targeted one
	time.Sleep(15 * time.Second)

	var abTestID int
	now := time.Now()

	abTest := analytics.ABTest{
		Name: cts.GenerateCanonicalPrefixName(),
		Variants: []analytics.Variant{
			{Index: indexName1, TrafficPercentage: 60, Description: "a description"},
			{Index: indexName2, TrafficPercentage: 40},
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

	// Find the AB test among all the existing AB tests and check it
	// corresponds to the original one
	{
		found := false

		res, err := analyticsClient.GetABTests(opt.Limit(100))
		require.NoError(t, err)
		for _, b := range res.ABTests {
			if b.ABTestID == abTestID {
				found = true
				checkABTestsAreEqual(t, abTest, b)
				break
			}
		}

		require.True(t, found)
	}

	// Stop the AB test
	{
		res, err := analyticsClient.StopABTest(abTestID)
		require.NoError(t, err)
		require.NoError(t, res.Wait())
	}

	// Check the AB test still exists but is stopped
	{
		res, err := analyticsClient.GetABTest(abTestID)
		require.NoError(t, err)
		require.Equal(t, res.Status, "stopped")
	}

	// Delete the AB test
	{
		res, err := analyticsClient.DeleteABTest(abTestID)
		require.NoError(t, err)
		require.NoError(t, res.Wait())
	}

	// Check the AB test doesn't exist anymore
	{
		_, err := analyticsClient.GetABTest(abTestID)
		require.Error(t, err)
	}
}
