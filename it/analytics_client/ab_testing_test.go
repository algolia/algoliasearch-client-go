package analytics_client

import (
	"strings"
	"testing"
	"time"

	"github.com/algolia/algoliasearch-client-go/algoliasearch"

	"github.com/stretchr/testify/require"

	"github.com/algolia/algoliasearch-client-go/it"
)

func TestABTesting(t *testing.T) {
	t.Parallel()
	client, index1, indexName1 := it.InitSearchClient1AndIndex(t)
	indexName2 := indexName1 + "_dev"
	index2 := client.InitIndex(indexName2)
	analytics := client.InitAnalytics()

	// Remove old AB tests
	{
		var toRemove []int
		today := it.TodayDate()

		res, err := analytics.GetABTests(nil)
		require.NoError(t, err)
		for _, abTest := range res.ABTests {
			if strings.HasPrefix("go-", abTest.Name) &&
				!strings.HasPrefix("go-"+today, abTest.Name) {
				toRemove = append(toRemove, abTest.ABTestID)
			}
		}

		for _, id := range toRemove {
			_, _ = analytics.DeleteABTest(id)
		}
	}

	// Create the two indices by adding a dummy object in each of them
	{
		res, err := index1.AddObject(algoliasearch.Object{"objectID": "one"})
		require.NoError(t, err)
		it.WaitTasks(t, index1, res.TaskID)

		res, err = index2.AddObject(algoliasearch.Object{"objectID": "one"})
		require.NoError(t, err)
		it.WaitTasks(t, index2, res.TaskID)
	}

	var abTestID int
	now := time.Now()

	abTest := algoliasearch.ABTest{
		Name: it.GenerateCanonicalPrefixName(),
		Variants: []algoliasearch.Variant{
			{Index: indexName1, TrafficPercentage: 60, Description: "a description"},
			{Index: indexName2, TrafficPercentage: 40},
		},
		EndAt: now.Truncate(time.Hour).Add(24 * time.Hour),
	}

	// Create the AB test
	{
		res, err := analytics.AddABTest(abTest)
		require.NoError(t, err)
		abTestID = res.ABTestID
		err = analytics.WaitTask(res)
		require.NoError(t, err)
	}

	// Retrieve the AB test and check it corresponds to the original one
	{
		res, err := analytics.GetABTest(abTestID)
		require.NoError(t, err)
		checkABTestsAreEqual(t, abTest, res)
		require.NotEqual(t, res.Status, "stopped")
	}

	// Find the AB test among all the existing AB tests and check it
	// corresponds to the original one
	{
		found := false

		res, err := analytics.GetABTests(nil)
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
		res, err := analytics.StopABTest(abTestID)
		require.NoError(t, err)
		err = analytics.WaitTask(res)
		require.NoError(t, err)
	}

	// Check the AB test still exists but is stopped
	{
		res, err := analytics.GetABTest(abTestID)
		require.NoError(t, err)
		require.Equal(t, res.Status, "stopped")
	}

	// Delete the AB test
	{
		res, err := analytics.DeleteABTest(abTestID)
		require.NoError(t, err)
		err = analytics.WaitTask(res)
		require.NoError(t, err)
	}

	// Check the AB test doesn't exist anymore
	{
		_, err := analytics.GetABTest(abTestID)
		require.Error(t, err)
	}
}

func checkABTestsAreEqual(t *testing.T, a algoliasearch.ABTest, b algoliasearch.ABTestResponse) {
	require.Equal(t, a.Name, b.Name)
	require.Equal(t, a.EndAt.Unix(), b.EndAt.Unix())
	require.Equal(t, len(a.Variants), len(b.Variants))

	var responseVariants []algoliasearch.Variant
	for _, v := range b.Variants {
		responseVariants = append(responseVariants, algoliasearch.Variant{
			Index:             v.Index,
			TrafficPercentage: v.TrafficPercentage,
			Description:       v.Description,
		})
	}
	require.ElementsMatch(t, a.Variants, responseVariants)
}
