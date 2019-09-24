package index

import (
	"io"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/opt"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/search"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/wait"
	"github.com/algolia/algoliasearch-client-go/v3/cts"
)

func TestQueryRule(t *testing.T) {
	t.Parallel()
	_, index, _ := cts.InitSearchClient1AndIndex(t)

	{
		res, err := index.SaveObjects([]map[string]string{
			{"objectID": "iphone_7", "brand": "Apple", "model": "7"},
			{"objectID": "iphone_8", "brand": "Apple", "model": "8"},
			{"objectID": "iphone_x", "brand": "Apple", "model": "X"},
			{"objectID": "one_plus_one", "brand": "OnePlus", "model": "One"},
			{"objectID": "one_plus_two", "brand": "OnePlus", "model": "Two"},
		})
		require.NoError(t, err)
		require.NoError(t, res.Wait())
	}

	g := wait.NewGroup()

	{
		res, err := index.SetSettings(search.Settings{
			AttributesForFaceting: opt.AttributesForFaceting("brand"),
		})
		require.NoError(t, err)
		g.Collect(res)
	}

	timeranges := []search.TimeRange{
		{
			From:  time.Date(2018, time.July, 24, 13, 35, 0, 0, time.UTC),
			Until: time.Date(2018, time.July, 25, 13, 35, 0, 0, time.UTC),
		},
		{
			From:  time.Date(2018, time.July, 26, 13, 35, 0, 0, time.UTC),
			Until: time.Date(2018, time.July, 27, 13, 35, 0, 0, time.UTC),
		},
	}

	rules := []search.Rule{
		{
			ObjectID:  "brand_automatic_faceting",
			Condition: search.RuleCondition{Anchoring: search.Contains, Pattern: "{facet:brand}"},
			Consequence: search.RuleConsequence{
				Params: &search.RuleParams{
					AutomaticFacetFilters: []search.AutomaticFacetFilter{
						{Facet: "brand", Disjunctive: true, Score: 42},
					},
				},
			},
			Validity:    timeranges,
			Description: "Automatic apply the faceting on `brand` if a brand value is found in the query",
			Enabled:     opt.Enabled(false),
		},
		{
			ObjectID: "query_edits",
			Condition: search.RuleCondition{
				Anchoring:    search.Contains,
				Pattern:      "mobile phone",
				Alternatives: search.AlternativesEnabled(),
			},
			Consequence: search.RuleConsequence{
				Params: &search.RuleParams{
					Query: search.NewRuleQueryObject(search.RuleQueryObjectQuery{
						Edits: []search.QueryEdit{
							search.RemoveEdit("mobile"),
							search.ReplaceEdit("phone", "iphone"),
						},
					}),
				},
			},
			Enabled: opt.Enabled(true),
		},
	}

	{
		res, err := index.SaveRule(rules[0])
		require.NoError(t, err)
		g.Collect(res)
	}

	{
		res, err := index.SaveRules(rules[1:])
		require.NoError(t, err)
		g.Collect(res)
	}

	require.NoError(t, g.Wait())

	{
		var wg sync.WaitGroup

		for _, rule := range rules {
			wg.Add(1)
			expected := rule
			go func(wg *sync.WaitGroup, expected search.Rule) {
				defer wg.Done()
				found, err := index.GetRule(expected.ObjectID)
				require.NoError(t, err, "should find rule whose objectID is %s", expected.ObjectID)
				require.True(t, found.Equal(expected))
			}(&wg, expected)
		}

		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			res, err := index.SearchRules("")
			require.NoError(t, err)

			found, err := res.Rules()
			require.NoError(t, err)

			checkRulesMatch(t, rules, found)
		}(&wg)

		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()

			var found []search.Rule
			it, err := index.BrowseRules()
			require.NoError(t, err)

			for {
				rule, err := it.Next()
				if err != nil {
					require.Equal(t, io.EOF, err)
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
		require.NoError(t, res.Wait())

		_, err = index.GetRule(rules[0].ObjectID)
		require.Error(t, err)
	}

	{
		res, err := index.ClearRules(false)
		require.NoError(t, err)
		require.NoError(t, res.Wait())
	}

	{
		res, err := index.SearchRules("")
		require.NoError(t, err)

		found, err := res.Rules()
		require.NoError(t, err)
		require.Len(t, found, 0)
	}
}

func checkRulesMatch(t *testing.T, expected, found []search.Rule) {
	require.Equal(t, len(expected), len(found))

	var count int

	for _, r1 := range expected {
		for _, r2 := range found {
			if r1.ObjectID == r2.ObjectID {
				require.True(t, r1.Equal(r2))
				count++
			}
		}
	}

	require.Equal(t, count, len(expected))
}
