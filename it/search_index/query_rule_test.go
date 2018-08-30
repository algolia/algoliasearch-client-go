package search_index

import (
	"sync"
	"testing"
	"time"

	"github.com/algolia/algoliasearch-client-go/algoliasearch"
	"github.com/algolia/algoliasearch-client-go/it"
	"github.com/stretchr/testify/require"
)

func TestQueryRule(t *testing.T) {
	t.Parallel()
	_, index, _ := it.InitSearchClient1AndIndex(t)

	{
		res, err := index.AddObjects([]algoliasearch.Object{
			{"objectID": "iphone_7", "brand": "Apple", "model": "7"},
			{"objectID": "iphone_8", "brand": "Apple", "model": "8"},
			{"objectID": "iphone_x", "brand": "Apple", "model": "X"},
			{"objectID": "one_plus_one", "brand": "OnePlus", "model": "One"},
			{"objectID": "one_plus_two", "brand": "OnePlus", "model": "Two"},
		})
		require.NoError(t, err)
		it.WaitTasks(t, index, res.TaskID)
	}

	var taskIDs []int

	{
		res, err := index.SetSettings(algoliasearch.Map{
			"attributesForFaceting": []string{"brand"},
		})
		require.NoError(t, err)
		taskIDs = append(taskIDs, res.TaskID)
	}

	timeranges := []algoliasearch.TimeRange{
		{
			From:  time.Date(2018, time.July, 24, 13, 35, 0, 0, time.UTC),
			Until: time.Date(2018, time.July, 25, 13, 35, 0, 0, time.UTC),
		},
		{
			From:  time.Date(2018, time.July, 26, 13, 35, 0, 0, time.UTC),
			Until: time.Date(2018, time.July, 27, 13, 35, 0, 0, time.UTC),
		},
	}

	rules := []algoliasearch.Rule{
		{
			ObjectID:  "brand_automatic_faceting",
			Condition: algoliasearch.NewSimpleRuleCondition(algoliasearch.Contains, "{facet:brand}"),
			Consequence: algoliasearch.RuleConsequence{
				Params: algoliasearch.Map{
					"automaticFacetFilters": []algoliasearch.AutomaticFacetFilter{
						{Facet: "brand", Disjunctive: true, Score: 42},
					},
				},
			},
			Validity:    timeranges,
			Description: "Automatic apply the faceting on `brand` if a brand value is found in the query",
		},
		{
			ObjectID:  "query_edits",
			Condition: algoliasearch.NewSimpleRuleCondition(algoliasearch.Contains, "mobile phone"),
			Consequence: algoliasearch.RuleConsequence{
				Params: algoliasearch.Map{
					"query": algoliasearch.Map{
						"edits": []algoliasearch.Edit{
							algoliasearch.DeleteEdit("mobile"),
							algoliasearch.ReplaceEdit("phone", "iphone"),
						},
					},
				},
			},
		},
	}

	rules[0].Disable()

	{
		res, err := index.SaveRule(rules[0], false)
		require.NoError(t, err)
		taskIDs = append(taskIDs, res.TaskID)
	}

	{
		res, err := index.BatchRules(rules[1:], false, false)
		require.NoError(t, err)
		taskIDs = append(taskIDs, res.TaskID)
	}

	it.WaitTasks(t, index, taskIDs...)
	taskIDs = []int{}

	{
		var wg sync.WaitGroup

		for _, rule := range rules {
			wg.Add(1)
			expected := rule
			objectID := rule.ObjectID
			go func(wg *sync.WaitGroup, objectID string, expected algoliasearch.Rule) {
				defer wg.Done()
				found, err := index.GetRule(objectID)
				require.NoError(t, err, "should find rule whose objectID is %s", objectID)
				require.Equal(t, expected.ObjectID, found.ObjectID)
			}(&wg, objectID, expected)
		}

		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			res, err := index.SearchRules(algoliasearch.Map{"query": ""})
			require.NoError(t, err)
			checkRulesMatch(t, rules, res.Hits)
		}(&wg)

		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			var found []algoliasearch.Rule
			it := algoliasearch.NewRuleIterator(index)

			for {
				rule, err := it.Next()
				if err != nil {
					require.Equal(t, algoliasearch.NoMoreRulesErr, err)
					break
				}
				found = append(found, *rule)
			}
			checkRulesMatch(t, rules, found)
		}(&wg)

		wg.Wait()
	}

	{
		res, err := index.DeleteRule(rules[0].ObjectID, false)
		require.NoError(t, err)
		it.WaitTasks(t, index, res.TaskID)

		_, err = index.GetRule(rules[0].ObjectID)
		require.Error(t, err)
	}

	{
		res, err := index.ClearRules(false)
		require.NoError(t, err)
		it.WaitTasks(t, index, res.TaskID)
	}

	{
		res, err := index.SearchRules(algoliasearch.Map{"query": ""})
		require.NoError(t, err)
		require.Len(t, res.Hits, 0)
	}
}

func checkRulesMatch(t *testing.T, expected, found []algoliasearch.Rule) {
	var expectedObjectIDs, foundObjectIDs []string

	for _, r := range expected {
		expectedObjectIDs = append(expectedObjectIDs, r.ObjectID)
	}

	for _, r := range found {
		foundObjectIDs = append(foundObjectIDs, r.ObjectID)
	}

	require.ElementsMatch(t, expectedObjectIDs, foundObjectIDs)
}
