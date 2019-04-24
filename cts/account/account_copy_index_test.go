package account

import (
	"testing"

	"github.com/algolia/algoliasearch-client-go/algolia"
	"github.com/algolia/algoliasearch-client-go/algolia/errs"
	"github.com/algolia/algoliasearch-client-go/algolia/opt"
	"github.com/algolia/algoliasearch-client-go/algolia/search"
	"github.com/algolia/algoliasearch-client-go/cts"
	"github.com/stretchr/testify/require"
)

func TestAccountCopyIndex(t *testing.T) {
	t.Parallel()
	client1, index1, indexName1 := cts.InitSearchClient1AndIndex(t)
	indexName2 := indexName1 + "_copy"
	index2 := client1.InitIndex(indexName2)

	{
		account := search.NewAccount()
		_, err := account.CopyIndex(index1, index2)
		require.Equal(t, errs.ErrSameAppID, err)
	}

	index2 = cts.InitSearchClient2(t).InitIndex(indexName2)
	await := algolia.Await()

	{
		res, err := index1.SaveObject(map[string]string{"objectID": "one"})
		require.NoError(t, err)
		await.Collect(res)
	}

	{
		res, err := index1.SaveRule(search.Rule{
			ObjectID:  "one",
			Condition: search.RuleCondition{Anchoring: search.Contains, Pattern: "pattern"},
			Consequence: search.RuleConsequence{
				Params: &search.RuleParams{
					Query: search.NewRuleQueryObject(search.RuleQueryObjectQuery{
						Edits: []search.QueryEdit{
							search.RemoveEdit("pattern"),
						},
					}),
				},
			},
		})
		require.NoError(t, err)
		await.Collect(res)
	}

	{
		res, err := index1.SaveSynonym(search.NewRegularSynonym("one", "one", "two"))
		require.NoError(t, err)
		await.Collect(res)
	}

	settings := search.Settings{
		SearchableAttributes: opt.SearchableAttributes("objectID"),
	}

	{
		res, err := index1.SetSettings(settings)
		require.NoError(t, err)
		await.Collect(res)
	}

	require.NoError(t, await.Wait())

	{
		account := search.NewAccount()
		wait, err := account.CopyIndex(index1, index2)
		require.NoError(t, err)
		require.NoError(t, wait.Wait())
	}

	{
		err := index2.GetObject("one", nil)
		require.NoError(t, err)

		_, err = index2.GetRule("one")
		require.NoError(t, err)

		_, err = index2.GetSynonym("one")
		require.NoError(t, err)

		found, err := index2.GetSettings()
		require.NoError(t, err)
		require.True(t, settings.Equal(found))
	}

	{
		account := search.NewAccount()
		_, err := account.CopyIndex(index1, index2)
		require.Equal(t, errs.ErrIndexAlreadyExists, err)
	}
}
