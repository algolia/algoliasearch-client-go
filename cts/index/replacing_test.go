package index

import (
	"testing"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/opt"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/search"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/wait"
	"github.com/algolia/algoliasearch-client-go/v3/cts"
	"github.com/stretchr/testify/require"
)

func TestReplacing(t *testing.T) {
	t.Parallel()
	_, index, _ := cts.InitSearchClient1AndIndex(t)

	g := wait.NewGroup()

	{
		res, err := index.SaveObject(map[string]string{"objectID": "one"})
		require.NoError(t, err)
		g.Collect(res)
	}

	{
		res, err := index.SaveRule(search.Rule{
			ObjectID:  "one",
			Condition: search.RuleCondition{Anchoring: search.Contains, Pattern: "pattern"},
			Consequence: search.RuleConsequence{
				Params: &search.RuleParams{
					Query: search.NewRuleQueryObject(
						search.RuleQueryObjectQuery{
							Edits: []search.QueryEdit{
								search.RemoveEdit("pattern"),
							},
						},
					),
				},
			},
		})
		require.NoError(t, err)
		g.Collect(res)
	}

	{
		res, err := index.SaveSynonym(search.NewRegularSynonym("one", "one", "two"))
		require.NoError(t, err)
		g.Collect(res)
	}

	require.NoError(t, g.Wait())

	{
		_, err := index.ReplaceAllObjects([]map[string]string{{"objectID": "two"}}, opt.Safe(true))
		require.NoError(t, err)
	}

	{
		res, err := index.ReplaceAllRules([]search.Rule{
			{
				ObjectID:  "two",
				Condition: search.RuleCondition{Anchoring: search.Contains, Pattern: "pattern"},
				Consequence: search.RuleConsequence{
					Params: &search.RuleParams{
						Query: search.NewRuleQueryObject(
							search.RuleQueryObjectQuery{
								Edits: []search.QueryEdit{
									search.RemoveEdit("pattern"),
								},
							},
						),
					},
				},
			},
		})
		require.NoError(t, err)
		g.Collect(res)
	}

	{
		res, err := index.ReplaceAllSynonyms([]search.Synonym{
			search.NewRegularSynonym("two", "one", "two"),
		})
		require.NoError(t, err)
		g.Collect(res)
	}

	require.NoError(t, g.Wait())

	{
		err := index.GetObject("one", nil)
		require.Error(t, err)

		err = index.GetObject("two", nil)
		require.NoError(t, err)

		_, err = index.GetRule("one")
		require.Error(t, err)

		_, err = index.GetRule("two")
		require.NoError(t, err)

		_, err = index.GetSynonym("one")
		require.Error(t, err)

		_, err = index.GetSynonym("two")
		require.NoError(t, err)
	}
}

func TestReplacingNonExistentIndex(t *testing.T) {
	t.Parallel()
	_, index, _ := cts.InitSearchClient1AndIndex(t)

	_, err := index.ReplaceAllObjects([]map[string]string{{"objectID": "one"}, {"objectID": "two"}}, opt.Safe(true))
	require.NoError(t, err)

	err = index.GetObject("one", nil)
	require.NoError(t, err)

	err = index.GetObject("two", nil)
	require.NoError(t, err)
}
