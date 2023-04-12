package index

import (
	"io"
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
			AttributesForFaceting: opt.AttributesForFaceting("brand", "model"),
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
			Conditions: []search.RuleCondition{
				{
					Anchoring:    search.Contains,
					Pattern:      "mobile phone",
					Alternatives: search.AlternativesEnabled(),
				},
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
		{
			ObjectID: "query_promo",
			Consequence: search.RuleConsequence{
				Params: &search.RuleParams{
					QueryParams: search.QueryParams{
						Filters: opt.Filters("brand:OnePlus"),
					},
				},
			},
		},
		{
			ObjectID: "query_promo_summer",
			Condition: search.RuleCondition{
				Context: "summer",
			},
			Consequence: search.RuleConsequence{
				Params: &search.RuleParams{
					QueryParams: search.QueryParams{
						Filters: opt.Filters("model:One"),
					},
				},
			},
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
		res, err := index.Search("", opt.RuleContexts("summer"))
		require.NoError(t, err, "should find search results")
		require.Equal(t, 1, res.NbHits)
	}

	for _, rule := range rules {
		found, err := index.GetRule(rule.ObjectID)
		require.NoError(t, err, "should find rule whose objectID is %s", rule.ObjectID)
		require.True(t, found.Equal(rule))
	}

	{
		res, err := index.SearchRules("")
		require.NoError(t, err)

		found, err := res.Rules()
		require.NoError(t, err)

		checkRulesMatch(t, rules, found)
	}

	{
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
