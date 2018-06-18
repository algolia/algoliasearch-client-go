package algoliasearch

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestABTesting(t *testing.T) {
	t.Parallel()
	c, a := initClientAndAnalytics(t)

	t.Log("TestABTesting: Remove any pre-existing AB test")
	for {
		res, err := a.GetABTests(nil)
		require.NoError(t, err)
		if res.Count == 0 || res.Total == 0 {
			break
		}

		var tasks []ABTestTaskRes
		for _, abTest := range res.ABTests {
			task, err := a.DeleteABTest(abTest.ABTestID)
			require.NoError(t, err)
			tasks = append(tasks, task)
		}
		for _, task := range tasks {
			err = a.WaitTask(task)
			require.NoError(t, err)
		}
	}

	now := time.Now()

	indexNamePrefix := fmt.Sprintf("TestABTesting_%d_", now.Unix())
	indexName1 := indexNamePrefix + "index1"
	indexName2 := indexNamePrefix + "index2"

	t.Log("TestABTesting: Prepare indices")
	i1 := initIndex(t, c, indexName1)
	i2 := initIndex(t, c, indexName2)
	addOneObject(t, i1)
	addOneObject(t, i2)

	abTest := ABTest{
		Name: "abtest_name",
		Variants: []Variant{
			Variant{Index: indexName1, TrafficPercentage: 60, Description: "a description"},
			Variant{Index: indexName2, TrafficPercentage: 40},
		},
		EndAt: now.Truncate(time.Hour).Add(24 * time.Hour),
	}

	var id int

	t.Log("TestABTesting: Add one AB Test")
	{
		task, err := a.AddABTest(abTest)
		if err != nil {
			t.Fatalf("TestABTesting: cannot add ab test: %s", err)
		}
		id = task.ABTestID

		t.Log("TestABTesting: Wait for AB test to be added")
		a.WaitTask(task)
	}

	t.Log("TestABTesting: Retrieve added AB test from all AB tests")
	{
		res, err := a.GetABTests(Map{
			"offset": 0,
			"limit":  10,
		})
		require.NoError(t, err)
		require.Equal(t, 1, res.Count)
		require.Equal(t, 1, res.Total)
		compareAbTests(t, id, abTest, res.ABTests[0])
	}

	t.Log("TestABTesting: Retrieve added AB test by ID")
	{
		res, err := a.GetABTest(id)
		require.NoError(t, err)
		compareAbTests(t, id, abTest, res)
	}

	t.Log("TestABTesting: Stop AB test")
	{
		task, err := a.StopABTest(id)
		require.NoError(t, err)
		a.WaitTask(task)

		res, err := a.GetABTest(id)
		require.NoError(t, err)
		require.Equal(t, "stopped", res.Status)
	}

	t.Log("TestABTesting: Delete AB test")
	{
		task, err := a.DeleteABTest(id)
		require.NoError(t, err)
		a.WaitTask(task)

		_, err = a.GetABTest(id)
		require.Error(t, err)
	}
}

func compareAbTests(t *testing.T, expectedID int, expectedABTest ABTest, got ABTestResponse) {
	require.Equal(t, expectedID, got.ABTestID)
	require.Equal(t, expectedABTest.Name, got.Name)
	require.True(t, expectedABTest.EndAt.Equal(got.EndAt), "was expecting %s but got %s", expectedABTest.EndAt, got.EndAt)
	compareVariants(t, expectedABTest.Variants, got.Variants)
}

func compareVariants(t *testing.T, expected []Variant, got []VariantResponse) {
	require.Equal(t, len(expected), len(got))

	found := 0
	for _, exp := range expected {
		for _, g := range got {
			if exp.Index == g.Index &&
				exp.TrafficPercentage == g.TrafficPercentage &&
				exp.Description == g.Description {
				found++
			}
		}
	}

	require.Equal(t, len(expected), found)
}
