package analytics_client

import (
	"strings"
	"testing"
	"time"

	"github.com/algolia/algoliasearch-client-go/algolia"
	"github.com/algolia/algoliasearch-client-go/algolia/analytics"
	"github.com/algolia/algoliasearch-client-go/cts"
	"github.com/stretchr/testify/require"
)

func TestABTesting(t *testing.T) {
	t.Parallel()
	searchClient, index1, indexName1 := cts.InitSearchClient1AndIndex(t)
	indexName2 := indexName1 + "_dev"
	index2 := searchClient.InitIndex(indexName2)
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

	// Create the two indices by adding a dummy object in each of them
	{
		await := algolia.Await()

		res, err := index1.SaveObject(map[string]string{"objectID": "one"})
		require.NoError(t, err)
		await.Collect(res)

		res, err = index2.SaveObject(map[string]string{"objectID": "one"})
		require.NoError(t, err)
		await.Collect(res)

		require.NoError(t, await.Wait())
	}

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

		res, err := analyticsClient.GetABTests()
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

func checkABTestsAreEqual(t *testing.T, a analytics.ABTest, b analytics.ABTestResponse) {
	require.Equal(t, a.Name, b.Name)
	require.Equal(t, a.EndAt.Unix(), b.EndAt.Unix())
	require.Equal(t, len(a.Variants), len(b.Variants))

	var responseVariants []analytics.Variant
	for _, v := range b.Variants {
		responseVariants = append(responseVariants, analytics.Variant{
			Index:             v.Index,
			TrafficPercentage: v.TrafficPercentage,
			Description:       v.Description,
		})
	}
	require.ElementsMatch(t, a.Variants, responseVariants)
}
